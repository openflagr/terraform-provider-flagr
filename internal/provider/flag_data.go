package provider

import (
	"context"
	"fmt"
	"strconv"
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
				Type:     schema.TypeString,
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

func dataSourceFlagRead(ctx context.Context, d *schema.ResourceData, i interface{}) (dgs diag.Diagnostics) {
	client := i.(*flagr.APIClient)
	errMsg := fmt.Sprintf("Unable to read flag %s", d.Get("id"))

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return Prettify(dgs, errMsg, err, false)
	}

	flag, _, err := client.FlagApi.GetFlag(context.TODO(), int64(id))
	if err != nil {
		return Prettify(dgs, errMsg, err, true)
	}

	m := map[string]interface{}{
		"id":                   d.Id(),
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
			return Prettify(dgs, errMsg, err, false)
		}
	}

	return dgs
}
