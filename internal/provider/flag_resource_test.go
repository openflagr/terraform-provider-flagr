package provider

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccCheckFlagDestroy(s *terraform.State) error {
	flags, _, err := APIClient().FlagApi.FindFlags(context.TODO(), nil)
	if err != nil {
		return fmt.Errorf("Cannot list flags: %w", err)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "flagr_flag" {
			continue
		}

		if len(flags) > 0 {
			for _, flag := range flags {
				if strconv.FormatInt(flag.Id, 10) == rs.Primary.ID {
					return fmt.Errorf("flag (%s) still exists", rs.Primary.ID)
				}
			}
		}
	}

	return nil
}

func testAccCheckFlagExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Missing Flagr ID")
		}

		id, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		_, _, err = APIClient().FlagApi.GetFlag(context.TODO(), int64(id))
		if err != nil {
			return fmt.Errorf("cannot find flag %s: %w", rs.Primary.ID, err)
		}

		return nil
	}
}

func TestAccContactGroup_basic(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		CheckDestroy:      testAccCheckFlagDestroy,
		ProviderFactories: providerFactory,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flagr_flag" "test" {
						description = "[TEST] Basic"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlagExists("flagr_flag.test"),
					resource.TestCheckResourceAttr("flagr_flag.test", "description", "[TEST] Basic"),
				),
			},
		},
	})
}
