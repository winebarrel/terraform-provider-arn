# arn:aws:forecast:ap-northeast-1:111111111111:explainability/resource-id
output "forecast_explainability" {
  value = provider::arn::forecast_explainability("resource-id")
}
