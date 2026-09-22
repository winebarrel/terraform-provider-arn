# arn:aws:iotdeviceadvisor:ap-northeast-1:111111111111:suitedefinition/suite-definition-id
output "iotdeviceadvisor_suitedefinition" {
  value = provider::arn::iotdeviceadvisor_suitedefinition("suite-definition-id")
}
