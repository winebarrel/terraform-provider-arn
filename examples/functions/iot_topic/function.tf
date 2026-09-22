# arn:aws:iot:ap-northeast-1:111111111111:topic/topic-name
output "iot_topic" {
  value = provider::arn::iot_topic("topic-name")
}
