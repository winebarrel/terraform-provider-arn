# arn:aws:quicksight:ap-northeast-1:111111111111:user/resource-id
output "quicksight_user" {
  value = provider::arn::quicksight_user("resource-id")
}
