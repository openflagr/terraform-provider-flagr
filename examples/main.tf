terraform {
  required_version = ">= 1.0.0"
  required_providers {
    flagr = {
      source  = "marceloboeira/flagr"
      version = "1.0.0"
    }
  }
}

module "ch" {
  source = "./pole"

  flag_name = "Switzerland"
}

output "ch" {
  value = module.ch.filtered_flags
}
