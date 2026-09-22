# arn:aws:quicksight:ap-northeast-1:111111111111:datasource/resource-id
output "quicksight_datasource" {
  value = provider::arn::quicksight_datasource("resource-id")
}
