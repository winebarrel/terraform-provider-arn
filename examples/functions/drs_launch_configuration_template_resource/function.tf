# arn:aws:drs:ap-northeast-1:111111111111:launch-configuration-template/launch-configuration-template-id
output "drs_launch_configuration_template_resource" {
  value = provider::arn::drs_launch_configuration_template_resource("launch-configuration-template-id")
}
