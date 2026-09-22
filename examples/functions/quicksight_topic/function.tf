# arn:aws:quicksight:ap-northeast-1:111111111111:topic/resource-id
output "quicksight_topic" {
  value = provider::arn::quicksight_topic("resource-id")
}
