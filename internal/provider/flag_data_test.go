package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccFlagData_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		CheckDestroy:      testAccCheckFlagDestroy,
		ProviderFactories: providerFactory,
		Steps: []resource.TestStep{
			{
				Config: `
					# Create a Flag
					resource "flagr_flag" "test" {
						description = "[TEST] Basic"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("flagr_flag.test", "id"),
				),
			},
			{
				Config: `
					resource "flagr_flag" "test" {
						description = "[TEST] Basic - Update 1"
					}

					# Use a Data Attribute to Fetch the Flag
					data "flagr_flag" "testd" {
						id = flagr_flag.test.key
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					//resource.TestCheckResourceAttrSet("flagr_flag.testd", "id"),
					resource.TestCheckResourceAttrSet("flagr_flag.testd", "description"),
				),
			},
		},
	})
}
