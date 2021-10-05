################### RESOURCE ###################

### Create
resource "flagr_flag" "li" {
  key         = "this-must-be-unique"
  description = "Lichtenstein"

  ### Modify
  # description = "Lichtenstein Update"
}


resource "flagr_flag" "fi" {
  enabled     = true
  description = "Finland"
  ### Uniqueness
  # key = "this-must-be-unique"
}

### Delete
# resource "flagr_flag" "br" {
#  description = "Brazil"
# }

output "countries" {
  value = [
    resource.flagr_flag.li.id,
    resource.flagr_flag.fi.id,
    # resource.flagr_flag.br.id,
  ]
}
