# arn:aws:cleanrooms-ml:ap-northeast-1:111111111111:configured-audience-model/resource-id
output "cleanrooms_ml_configuredaudiencemodel" {
  value = provider::arn::cleanrooms_ml_configuredaudiencemodel("resource-id")
}
