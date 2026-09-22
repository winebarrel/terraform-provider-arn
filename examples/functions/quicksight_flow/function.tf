# arn:aws:quicksight:ap-northeast-1:111111111111:flow/resource-id
output "quicksight_flow" {
  value = provider::arn::quicksight_flow("resource-id")
}
