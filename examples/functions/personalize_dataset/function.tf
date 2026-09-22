# arn:aws:personalize:ap-northeast-1:111111111111:dataset/resource-id
output "personalize_dataset" {
  value = provider::arn::personalize_dataset("resource-id")
}
