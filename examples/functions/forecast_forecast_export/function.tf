# arn:aws:forecast:ap-northeast-1:111111111111:forecast-export-job/resource-id
output "forecast_forecast_export" {
  value = provider::arn::forecast_forecast_export("resource-id")
}
