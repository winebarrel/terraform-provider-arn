# arn:aws:quicksight:ap-northeast-1:111111111111:group/resource-id
output "quicksight_group" {
  value = provider::arn::quicksight_group("resource-id")
}
