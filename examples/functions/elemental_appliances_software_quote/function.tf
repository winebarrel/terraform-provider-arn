# arn:aws:elemental-appliances-software:ap-northeast-1:111111111111:quote/resource-id
output "elemental_appliances_software_quote" {
  value = provider::arn::elemental_appliances_software_quote("resource-id")
}
