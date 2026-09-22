# arn:aws:one:ap-northeast-1:111111111111:device-configuration-template/template-id
output "one_device_configuration_template" {
  value = provider::arn::one_device_configuration_template("template-id")
}
