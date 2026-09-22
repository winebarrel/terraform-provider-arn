# arn:aws:appconfig:ap-northeast-1:111111111111:application/application-id/experimentdefinition/experiment-definition-id
output "appconfig_experimentdefinition" {
  value = provider::arn::appconfig_experimentdefinition("application-id", "experiment-definition-id")
}
