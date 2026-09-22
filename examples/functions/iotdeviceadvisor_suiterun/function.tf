# arn:aws:iotdeviceadvisor:ap-northeast-1:111111111111:suiterun/suite-definition-id/suite-run-id
output "iotdeviceadvisor_suiterun" {
  value = provider::arn::iotdeviceadvisor_suiterun("suite-definition-id", "suite-run-id")
}
