package provider

import (
	"encoding/json"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	flagr "github.com/openflagr/goflagr"
)

type errorResponse struct {
	Message string `json:"message"`
}

// ParseAPIError - Parse the API Error from swagger
func ParseAPIError(e error) string {
	var res errorResponse
	body := e.(flagr.GenericSwaggerError).Body()
	err := json.Unmarshal(body, &res)
	if err != nil {
		// Return the raw body when not JSON valid
		return string(body)
	}

	return res.Message
}

// Prettify - Inject terraform context into errors
func Prettify(dgs diag.Diagnostics, m string, e error, hydrate bool) diag.Diagnostics {
	d := e.Error()
	if hydrate {
		d = ParseAPIError(e)
	}
	return append(dgs, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  m,
		Detail:   d,
	})
}
