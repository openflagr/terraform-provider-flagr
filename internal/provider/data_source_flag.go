package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/openflagr/goflagr"
)

func dataSourceFlags() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFlagsRead,
		Schema: map[string]*schema.Schema{
			// https://github.com/openflagr/goflagr/blob/main/model_flag.go
			"flags": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
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
				},
			},
		},
	}
}

func dataSourceFlagsRead(ctx context.Context, d *schema.ResourceData, i interface{}) diag.Diagnostics {
	client := i.(*goflagr.APIClient)

	var diags diag.Diagnostics

	flags, _, err := client.FlagApi.FindFlags(context.TODO(), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	// TODO Move to its own method? FlagToTerraform?
	// Converting Flag Schema
	var pF []map[string]interface{}
	for _, flag := range flags {
		pF = append(pF,
			map[string]interface{}{
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
			},
		)
	}

	if err := d.Set("flags", pF); err != nil {
		// Improve error message: https://learn.hashicorp.com/tutorials/terraform/provider-debug?in=terraform/providers
		return diag.FromErr(err)
	}

	// always run ??
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
