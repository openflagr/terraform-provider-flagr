terraform {
  required_version = ">= 1.0.0"
  required_providers {
    flagr = {
      source = "marceloboeira/flagr"
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

output "filtered_flag" {
  value = {
    for flag in data.flagr_flags.all.flags :
    flag.id => flag
    if flag.enabled == false
  }
}
