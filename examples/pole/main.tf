terraform {
  required_version = ">= 1.0.0"
  required_providers {
    flagr = {
      source  = "marceloboeira/flagr"
      version = "1.0.0"
    }
  }
}

variable "flag_name" {
  type    = string
  default = "Switzerland"
}

data "flagr_flags" "all" {}

output "all_flags" {
  value = data.flagr_flags.all.flags
}

output "flag" {
  value = {
    for flag in data.flagr_flags.all.flags :
    flag.id => flag
    if flag.description == var.flag_name
  }
}
