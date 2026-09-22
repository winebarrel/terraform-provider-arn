# arn:aws:forecast:ap-northeast-1:111111111111:explainability-export/resource-id
output "forecast_explainability_export" {
  value = provider::arn::forecast_explainability_export("resource-id")
}
