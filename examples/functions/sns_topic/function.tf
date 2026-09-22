# arn:aws:sns:ap-northeast-1:111111111111:topic-name
output "sns_topic" {
  value = provider::arn::sns_topic("topic-name")
}
