# arn:aws:quicksight:ap-northeast-1:111111111111:custompermissions/resource-id
output "quicksight_custompermissions" {
  value = provider::arn::quicksight_custompermissions("resource-id")
}
