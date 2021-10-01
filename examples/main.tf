terraform {
  required_version = ">= 1.0.0"
  required_providers {
    flagr = {
      source  = "marceloboeira/flagr"
      version = "1.0.0"
    }
  }
}

provider "flagr" {
  host = "http://0.0.0.0:18000/"
  # path = "/api/v1" # Optional
}


// Expects 2 flags to work
// ID: 1 - Switzerland
// ID: 2 - Germany

// All flags
data "flagr_flags" "all" {}

// Specific Flag with ID
data "flagr_flag" "de" {
  id = 2
}

// Invalid Flag (Not found)
// data "flagr_flag" "invalid" {
//   id = 666
// }

output "ch" {
  value = element(data.flagr_flags.all.flags, 0).description
}

output "de" {
  value = data.flagr_flag.de.description
}
