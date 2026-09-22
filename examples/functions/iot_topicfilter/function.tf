# arn:aws:iot:ap-northeast-1:111111111111:topicfilter/topic-filter
output "iot_topicfilter" {
  value = provider::arn::iot_topicfilter("topic-filter")
}
