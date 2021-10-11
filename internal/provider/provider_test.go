package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	flagr "github.com/openflagr/goflagr"
)

var providerFactory = map[string]func() (*schema.Provider, error){
	"flagr": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}

func testAccPreCheck(t *testing.T) func() {
	return func() {
		t.Helper()

		if host := os.Getenv(FLAGR_HOST); host == "" {
			t.Fatal("Missing FLAGR_HOST")
		}

		if path := os.Getenv(FLAGR_PATH); path == "" {
			t.Fatal("Missing FLAGR_PATH")
		}
	}
}

func TestProvider(t *testing.T) {
	t.Parallel()

	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func APIClient() *flagr.APIClient {
	return newAPIClient(
		os.Getenv(FLAGR_HOST) + os.Getenv(FLAGR_PATH),
	)
}
