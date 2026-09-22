# arn:aws:forecast:ap-northeast-1:111111111111:monitor/resource-id
output "forecast_monitor" {
  value = provider::arn::forecast_monitor("resource-id")
}
