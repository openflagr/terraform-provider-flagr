################### RESOURCE ###################

resource "flagr_flag" "li" {
  # key = "lich"
  description = "Lichtenstein"
}

output "li" {
  value = resource.flagr_flag.li.id
}
