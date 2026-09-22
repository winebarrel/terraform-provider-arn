# arn:aws:forecast:ap-northeast-1:111111111111:dataset-import-job/resource-id
output "forecast_dataset_import_job" {
  value = provider::arn::forecast_dataset_import_job("resource-id")
}
