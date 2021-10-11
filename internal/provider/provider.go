package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	flagr "github.com/openflagr/goflagr"
)

// FLAGR_HOST - Name of ENV variable to look for flagr host
const FLAGR_HOST = "FLAGR_HOST"

// FLAGR_PATH - Name of ENV variable to look for flagr path
const FLAGR_PATH = "FLAGR_PATH"

// FLAGR_AGENT - Static HTTP Agent Name to query flagr instances
const FLAGR_AGENT = "Terraform"

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

	return newAPIClient(u.String()), dg
}

func newAPIClient(url string) *flagr.APIClient {
	return flagr.NewAPIClient(&flagr.Configuration{
		BasePath:  url,
		UserAgent: FLAGR_AGENT,
	})
}
