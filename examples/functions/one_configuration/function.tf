# arn:aws:one:ap-northeast-1:111111111111:device-instance/device-instance-id/configuration/version
output "one_configuration" {
  value = provider::arn::one_configuration("device-instance-id", "version")
}
