# arn:aws:forecast:ap-northeast-1:111111111111:dataset/resource-id
output "forecast_dataset" {
  value = provider::arn::forecast_dataset("resource-id")
}
