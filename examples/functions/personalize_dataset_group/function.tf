# arn:aws:personalize:ap-northeast-1:111111111111:dataset-group/resource-id
output "personalize_dataset_group" {
  value = provider::arn::personalize_dataset_group("resource-id")
}
