package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const RFC3339 = `^((?:(\d{4}-\d{2}-\d{2})T(\d{2}:\d{2}:\d{2}(?:\.\d+)?))(Z|[\+-]\d{2}:\d{2})?)$`

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
					// Given attributes
					resource.TestCheckResourceAttr("flagr_flag.test", "description", "[TEST] Basic"),

					// Default attributes
					resource.TestCheckResourceAttr("flagr_flag.test", "enabled", "false"),
					resource.TestCheckResourceAttr("flagr_flag.test", "data_records_enabled", "false"),

					// Computed attributes
					resource.TestCheckResourceAttrSet("flagr_flag.test", "id"),
					resource.TestMatchResourceAttr("flagr_flag.test", "updated_at", regexp.MustCompile(RFC3339)),
				),
			},
			{
				Config: `
					resource "flagr_flag" "test" {
						description = "[TEST] Basic - Update 1"

						enabled              = true
						data_records_enabled = true
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlagExists("flagr_flag.test"),
					// Given attributes
					resource.TestCheckResourceAttr("flagr_flag.test", "description", "[TEST] Basic - Update 1"),
					resource.TestCheckResourceAttr("flagr_flag.test", "enabled", "true"),
					resource.TestCheckResourceAttr("flagr_flag.test", "data_records_enabled", "true"),

					// Computed attributes
					resource.TestCheckResourceAttrSet("flagr_flag.test", "id"),
					resource.TestMatchResourceAttr("flagr_flag.test", "updated_at", regexp.MustCompile(RFC3339)),
				),
			},
		},
	})
}

func TestAccContactGroup_complete(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		CheckDestroy:      testAccCheckFlagDestroy,
		ProviderFactories: providerFactory,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flagr_flag" "test" {
						description = "[TEST] Complete"

						enabled = true
						data_records_enabled = true
						notes = "Managed by Terraform"
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlagExists("flagr_flag.test"),
					// Given attributes
					resource.TestCheckResourceAttr("flagr_flag.test", "description", "[TEST] Complete"),
					resource.TestCheckResourceAttr("flagr_flag.test", "enabled", "true"),
					resource.TestCheckResourceAttr("flagr_flag.test", "data_records_enabled", "true"),
					resource.TestCheckResourceAttr("flagr_flag.test", "notes", "Managed by Terraform"),

					// Computed attributes
					resource.TestCheckResourceAttrSet("flagr_flag.test", "id"),
					resource.TestMatchResourceAttr("flagr_flag.test", "updated_at", regexp.MustCompile(RFC3339)),
				),
			},
		},
	})
}

func TestAccContactGroup_validations(t *testing.T) {
	t.Parallel()
	rKey := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:          testAccPreCheck(t),
		CheckDestroy:      testAccCheckFlagDestroy,
		ProviderFactories: providerFactory,
		Steps: []resource.TestStep{
			{
				Config: `
					resource "flagr_flag" "test" {
						description = ""
					}
				`,
				ExpectError: regexp.MustCompile("expected \"description\" to not be an empty string"),
			},
			{
				Config: fmt.Sprintf(`
					resource "flagr_flag" "test_a" {
						description = "[TEST] Duplicated Key 1"
						key = "%s"
					}

					resource "flagr_flag" "test_b" {
						description = "[TEST] Duplicated Key 2"
						key = "%s-not-dup"
					}

					resource "flagr_flag" "test_c" {
						description = "[TEST] Duplicated Key 3"
						key = "%s"
					}
				`, rKey, rKey, rKey),
				ExpectError: regexp.MustCompile("cannot create flag. UNIQUE constraint failed: flags.key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlagExists("flagr_flag.test_a"),
					testAccCheckFlagExists("flagr_flag.test_b"),
				),
			},
			{
				Config: fmt.Sprintf(`
					resource "flagr_flag" "test_a" {
						description = "[TEST] Duplicated Key 1"
						key = "%s"
					}

					resource "flagr_flag" "test_b" {
						description = "[TEST] Duplicated Key 2"
						key = "%s"
					}
				`, rKey, rKey),
				ExpectError: regexp.MustCompile("UNIQUE constraint failed: flags.key"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckFlagExists("flagr_flag.test_a"),
					testAccCheckFlagExists("flagr_flag.test_b"),
					resource.TestCheckResourceAttr("flagr_flag.test_a", "key", fmt.Sprintf("%s", rKey)),
					resource.TestCheckResourceAttr("flagr_flag.test_b", "key", fmt.Sprintf("%s-not-dup", rKey)),
				),
			},
		},
	})
}
