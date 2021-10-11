package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	flagr "github.com/openflagr/goflagr"
)

const FLAGR_HOST = "FLAGR_HOST"
const FLAGR_PATH = "FLAGR_PATH"

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(FLAGR_HOST, nil),
				Description: "Host for the flagr API, e.g.: flagr.mycompany.com",
			},
			"path": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc(FLAGR_PATH, "/api/v1"),
				Description: "Path for the flagr API, e.g.: /custom/api/v1, default: /api/v1",
			},
		},
		ConfigureContextFunc: providerConfigure,
		ResourcesMap: map[string]*schema.Resource{
			"flagr_flag": resourceFlag(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"flagr_flag":  dataSourceFlag(),
			"flagr_flags": dataSourceFlags(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (i interface{}, dg diag.Diagnostics) {
	rURL := d.Get("host").(string) + d.Get("path").(string)
	u, err := url.ParseRequestURI(rURL)
	if err != nil {
		// TODO Improve error message: https://learn.hashicorp.com/tutorials/terraform/provider-debug?in=terraform/providers
		return nil, diag.FromErr(fmt.Errorf("%s is not a valid URL", rURL))
	}

	c := flagr.NewAPIClient(&flagr.Configuration{
		BasePath:  u.String(),
		UserAgent: "Terraform",
	})

	return c, dg
}
