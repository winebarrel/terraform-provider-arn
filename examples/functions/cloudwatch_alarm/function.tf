# arn:aws:cloudwatch:ap-northeast-1:111111111111:alarm:alarm-name
output "cloudwatch_alarm" {
  value = provider::arn::cloudwatch_alarm("alarm-name")
}
