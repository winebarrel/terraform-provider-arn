# arn:aws:iot:ap-northeast-1:111111111111:scheduledaudit/schedule-name
output "iot_scheduledaudit" {
  value = provider::arn::iot_scheduledaudit("schedule-name")
}
