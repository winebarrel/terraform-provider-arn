# arn:aws:quicksight:ap-northeast-1:111111111111:dataset/resource-id
output "quicksight_dataset" {
  value = provider::arn::quicksight_dataset("resource-id")
}
