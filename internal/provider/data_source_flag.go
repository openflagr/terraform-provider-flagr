package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFlags() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFlagsRead,
		Schema: map[string]*schema.Schema{
			"flags": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						//					"data_records_enabled": &schema.Schema{
						//						Type:     schema.TypeBool,
						//						Computed: true,
						//					},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": &schema.Schema{
							Type:     schema.TypeBool,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"key": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"segments": &schema.Schema{
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString, // TODO FIX TYPE
							},
							Computed: true,
						},
						"tags": &schema.Schema{
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString, //TODO FIX TYPE
							},
							Computed: true,
						},
						//					"updated_at": &schema.Schema{
						//						Type:         schema.TypeString,
						//						ValidateFunc: validation.IsRFC3339Time,
						//						Computed:     true,
						//					},
						"variants": &schema.Schema{
							Type: schema.TypeList,
							Elem: &schema.Schema{
								Type: schema.TypeString, //TODO FIX TYPE
							},
							Computed: true,
						},

						// TODO ADD notes
					},
				},
			},
		},
	}
}

func dataSourceFlagsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/flags", "http://0.0.0.0:18000/api/v1"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	flags := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&flags)
	if err != nil {
		return diag.FromErr(err)
	}

	// Keys are currently not matching the expected schema
	for _, flag := range flags {
		delete(flag, "updatedAt")
		delete(flag, "dataRecordsEnabled")
		delete(flag, "notes")
	}

	if err := d.Set("flags", flags); err != nil {
		return diag.FromErr(err)
	}

	// always run ??
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
