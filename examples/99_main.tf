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
  host = "http://0.0.0.0:18000"
  path = "/api/v1" # Optional
  # TODO: Authentication
}
