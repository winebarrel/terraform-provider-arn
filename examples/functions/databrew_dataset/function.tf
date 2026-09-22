# arn:aws:databrew:ap-northeast-1:111111111111:dataset/resource-id
output "databrew_dataset" {
  value = provider::arn::databrew_dataset("resource-id")
}
