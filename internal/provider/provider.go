package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	flagr "github.com/openflagr/goflagr"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/v1",
			},
		},
		ConfigureContextFunc: providerConfigure,
		ResourcesMap:         map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"flagr_flags": dataSourceFlags(),
		},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	rURL := d.Get("host").(string) + d.Get("path").(string)
	u, err := url.ParseRequestURI(rURL)
	if err != nil {
		return nil, diag.FromErr(fmt.Errorf("%s is not a valid URL", rURL))
	}

	c := flagr.NewAPIClient(&flagr.Configuration{
		BasePath:  u.String(),
		UserAgent: "Terraform",
	})

	return c, diags
}
