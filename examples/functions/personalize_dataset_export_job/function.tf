# arn:aws:personalize:ap-northeast-1:111111111111:dataset-export-job/resource-id
output "personalize_dataset_export_job" {
  value = provider::arn::personalize_dataset_export_job("resource-id")
}
