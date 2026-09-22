# arn:aws:quicksight:ap-northeast-1:111111111111:app/resource-id
output "quicksight_app" {
  value = provider::arn::quicksight_app("resource-id")
}
