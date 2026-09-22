# arn:aws:personalize:ap-northeast-1:111111111111:data-deletion-job/resource-id
output "personalize_data_deletion_job" {
  value = provider::arn::personalize_data_deletion_job("resource-id")
}
