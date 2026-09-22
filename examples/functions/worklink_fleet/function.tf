# arn:aws:worklink::111111111111:fleet/fleet-name
output "worklink_fleet" {
  value = provider::arn::worklink_fleet("fleet-name")
}
