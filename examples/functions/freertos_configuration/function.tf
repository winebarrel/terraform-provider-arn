# arn:aws:freertos:ap-northeast-1:111111111111:configuration/configuration-name
output "freertos_configuration" {
  value = provider::arn::freertos_configuration("configuration-name")
}
