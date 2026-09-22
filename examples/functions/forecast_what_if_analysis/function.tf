# arn:aws:forecast:ap-northeast-1:111111111111:what-if-analysis/resource-id
output "forecast_what_if_analysis" {
  value = provider::arn::forecast_what_if_analysis("resource-id")
}
