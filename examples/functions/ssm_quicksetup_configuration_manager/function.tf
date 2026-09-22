# arn:aws:ssm-quicksetup:ap-northeast-1:111111111111:configuration-manager/configuration-manager-id
output "ssm_quicksetup_configuration_manager" {
  value = provider::arn::ssm_quicksetup_configuration_manager("configuration-manager-id")
}
