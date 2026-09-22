# arn:aws:personalize:ap-northeast-1:111111111111:dataset-import-job/resource-id
output "personalize_dataset_import_job" {
  value = provider::arn::personalize_dataset_import_job("resource-id")
}
