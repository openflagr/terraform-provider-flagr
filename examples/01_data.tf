################### DATA ###################
## Expects 2 flags to work
## ID: 1 - Switzerland
## ID: 2 - Germany

## All flags
data "flagr_flags" "all" {}

data "flagr_flag" "ch" {
  id = data.flagr_flags.all.flags[index(data.flagr_flags.all.flags.*.description, "Switzerland")].id
}

data "flagr_flag" "de" {
  id = data.flagr_flags.all.flags[index(data.flagr_flags.all.flags.*.description, "Germany")].id
}

##  Invalid Flag (Not found)
##  data "flagr_flag" "invalid" {
##    id = 666
##  }

output "ch" {
  value = data.flagr_flag.ch.id
}

output "de" {
  value = data.flagr_flag.de.id
}
