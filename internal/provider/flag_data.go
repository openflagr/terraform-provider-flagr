package provider

import (
	"context"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	flagr "github.com/openflagr/goflagr"
)

func dataSourceFlag() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFlagRead,
		// https://github.com/openflagr/goflagr/blob/main/model_flag.go
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"segments": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString, // TODO FIX TYPE
				},
				Computed: true,
			},
			"variants": &schema.Schema{
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString, //TODO FIX TYPE
				},
				Computed: true,
			},
			"data_records_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Computed: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceFlagRead(ctx context.Context, d *schema.ResourceData, i interface{}) (diags diag.Diagnostics) {
	client := i.(*flagr.APIClient)

	flag, _, err := client.FlagApi.GetFlag(context.TODO(), int64(d.Get("id").(int)))
	if err != nil {
		return diag.FromErr(err)
	}

	m := map[string]interface{}{
		"id":                   flag.Id,
		"key":                  flag.Key,
		"description":          flag.Description,
		"enabled":              flag.Enabled,
		"segments":             flag.Segments,
		"variants":             flag.Variants,
		"data_records_enabled": flag.DataRecordsEnabled,
		"notes":                flag.Notes,
		"created_by":           flag.CreatedBy,
		"updated_by":           flag.UpdatedBy,
		"updated_at":           flag.UpdatedAt.Format(time.RFC3339),
	}
	for k, v := range m {
		if err := d.Set(k, v); err != nil {
			// Improve error message: https://learn.hashicorp.com/tutorials/terraform/provider-debug?in=terraform/providers
			return diag.FromErr(err)
		}
	}

	return diags
}
