# arn:aws:quicksight:ap-northeast-1:111111111111:customization/resource-id
output "quicksight_customization" {
  value = provider::arn::quicksight_customization("resource-id")
}
