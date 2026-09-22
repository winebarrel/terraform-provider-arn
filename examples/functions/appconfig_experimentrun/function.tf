# arn:aws:appconfig:ap-northeast-1:111111111111:application/application-id/experimentdefinition/experiment-definition-id/experimentrun/experiment-run-number
output "appconfig_experimentrun" {
  value = provider::arn::appconfig_experimentrun("application-id", "experiment-definition-id", "experiment-run-number")
}
