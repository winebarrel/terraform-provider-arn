# arn:aws:quicksight:ap-northeast-1:111111111111:theme/resource-id
output "quicksight_theme" {
  value = provider::arn::quicksight_theme("resource-id")
}
