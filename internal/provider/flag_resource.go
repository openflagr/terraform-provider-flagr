package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/openflagr/goflagr"
)

func resourceFlag() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceFlagRead,
		CreateContext: resourceFlagCreate,
		UpdateContext: resourceFlagUpdate,
		DeleteContext: resourceFlagDelete,
		Schema: map[string]*schema.Schema{
			//"id": &schema.Schema{
			//	Type:     schema.TypeInt,
			//	Computed: true,
			//},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			// "segments": &schema.Schema{
			// 	Type: schema.TypeList,
			// 	Elem: &schema.Schema{
			// 		Type: schema.TypeString, // TODO FIX TYPE
			// 	},
			// 	Computed: true,
			// },
			//"variants": &schema.Schema{
			//	Type: schema.TypeList,
			//	Elem: &schema.Schema{
			//		Type: schema.TypeString, //TODO FIX TYPE
			//	},
			//	Computed: true,
			//},
			"data_records_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"updated_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"updated_at": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceFlagRead(ctx context.Context, d *schema.ResourceData, i interface{}) (dg diag.Diagnostics) {
	client := i.(*goflagr.APIClient)

	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		// TODO: Improve error message
		return diag.FromErr(err)
	}

	flag, _, err := client.FlagApi.GetFlag(context.TODO(), int64(id))
	if err != nil {
		return diag.FromErr(err)
	}

	m := map[string]interface{}{
		/////"id":          id,
		"key":         flag.Key,
		"description": flag.Description,
		"enabled":     flag.Enabled,
		//"segments":             flag.Segments,
		//"variants":             flag.Variants,
		"data_records_enabled": flag.DataRecordsEnabled,
		"notes":                flag.Notes,
		"created_by":           flag.CreatedBy,
		"updated_by":           flag.UpdatedBy,
		"updated_at":           flag.UpdatedAt.Format(time.RFC3339),
	}
	for k, v := range m {
		if err := d.Set(k, v); err != nil {
			// TODO Improve error message: https://learn.hashicorp.com/tutorials/terraform/provider-debug?in=terraform/providers
			return diag.FromErr(err)
		}
	}

	return dg
}

func resourceFlagCreate(ctx context.Context, d *schema.ResourceData, i interface{}) (dg diag.Diagnostics) {
	client := i.(*goflagr.APIClient)

	flag, _, err := client.FlagApi.CreateFlag(context.TODO(), goflagr.CreateFlagRequest{
		Description: d.Get("description").(string),
		Key:         d.Get("key").(string),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(flag.Id, 10))

	return resourceFlagUpdate(ctx, d, i)
}

func resourceFlagUpdate(ctx context.Context, d *schema.ResourceData, i interface{}) (dg diag.Diagnostics) {
	client := i.(*goflagr.APIClient)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		// TODO: Improve error message
		return diag.FromErr(err)
	}

	if d.HasChanges("key", "description", "data_records_enabled", "notes") {
		_, _, err = client.FlagApi.PutFlag(context.TODO(), id, goflagr.PutFlagRequest{
			Key:                d.Get("key").(string),
			Description:        d.Get("description").(string),
			DataRecordsEnabled: d.Get("data_records_enabled").(bool),
			Notes:              d.Get("notes").(string),
		})
		if err != nil {
			// TODO: Improve error message
			return diag.FromErr(err)
		}
	}

	if d.HasChange("enabled") {
		_, _, err = client.FlagApi.SetFlagEnabled(context.TODO(), id, goflagr.SetFlagEnabledRequest{
			Enabled: d.Get("enabled").(bool),
		})
		if err != nil {
			// TODO: Improve error message
			return diag.FromErr(err)
		}
	}

	return resourceFlagRead(ctx, d, i)
}

func resourceFlagDelete(ctx context.Context, d *schema.ResourceData, i interface{}) (dg diag.Diagnostics) {
	client := i.(*goflagr.APIClient)
	id, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}

	_, err = client.FlagApi.DeleteFlag(context.TODO(), id)
	if err != nil {
		return diag.FromErr(err)
	}

	return dg
}
