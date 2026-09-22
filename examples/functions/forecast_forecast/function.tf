# arn:aws:forecast:ap-northeast-1:111111111111:forecast/resource-id
output "forecast_forecast" {
  value = provider::arn::forecast_forecast("resource-id")
}
