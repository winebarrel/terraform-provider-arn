# arn:aws:forecast:ap-northeast-1:111111111111:forecast-endpoint/resource-id
output "forecast_endpoint" {
  value = provider::arn::forecast_endpoint("resource-id")
}
