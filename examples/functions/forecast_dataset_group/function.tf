# arn:aws:forecast:ap-northeast-1:111111111111:dataset-group/resource-id
output "forecast_dataset_group" {
  value = provider::arn::forecast_dataset_group("resource-id")
}
