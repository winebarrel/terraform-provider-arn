# arn:aws:quicksight:ap-northeast-1:111111111111:extension/resource-id
output "quicksight_extension" {
  value = provider::arn::quicksight_extension("resource-id")
}
