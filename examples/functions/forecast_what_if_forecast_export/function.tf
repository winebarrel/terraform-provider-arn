# arn:aws:forecast:ap-northeast-1:111111111111:what-if-forecast-export/resource-id
output "forecast_what_if_forecast_export" {
  value = provider::arn::forecast_what_if_forecast_export("resource-id")
}
