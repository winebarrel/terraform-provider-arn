# arn:aws:iot:ap-northeast-1:111111111111:provisioningtemplate/provisioning-template
output "iot_provisioningtemplate" {
  value = provider::arn::iot_provisioningtemplate("provisioning-template")
}
