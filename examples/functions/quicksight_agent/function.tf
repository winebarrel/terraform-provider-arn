# arn:aws:quicksight:ap-northeast-1:111111111111:agent/resource-id
output "quicksight_agent" {
  value = provider::arn::quicksight_agent("resource-id")
}
