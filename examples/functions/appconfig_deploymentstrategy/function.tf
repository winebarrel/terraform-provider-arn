# arn:aws:appconfig:ap-northeast-1:111111111111:deploymentstrategy/deployment-strategy-id
output "appconfig_deploymentstrategy" {
  value = provider::arn::appconfig_deploymentstrategy("deployment-strategy-id")
}
