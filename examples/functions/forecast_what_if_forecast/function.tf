# arn:aws:forecast:ap-northeast-1:111111111111:what-if-forecast/resource-id
output "forecast_what_if_forecast" {
  value = provider::arn::forecast_what_if_forecast("resource-id")
}
