# arn:aws:quicksight:ap-northeast-1:111111111111:space/resource-id
output "quicksight_space" {
  value = provider::arn::quicksight_space("resource-id")
}
