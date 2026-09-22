# arn:aws:forecast:ap-northeast-1:111111111111:predictor/resource-id
output "forecast_predictor" {
  value = provider::arn::forecast_predictor("resource-id")
}
