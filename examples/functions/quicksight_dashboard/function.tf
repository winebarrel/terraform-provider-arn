# arn:aws:quicksight:ap-northeast-1:111111111111:dashboard/resource-id
output "quicksight_dashboard" {
  value = provider::arn::quicksight_dashboard("resource-id")
}
