# arn:aws:quicksight:ap-northeast-1:111111111111:account/resource-id
output "quicksight_account" {
  value = provider::arn::quicksight_account("resource-id")
}
