# arn:aws:quicksight:ap-northeast-1:111111111111:brand/resource-id
output "quicksight_brand" {
  value = provider::arn::quicksight_brand("resource-id")
}
