# arn:aws:medialive:ap-northeast-1:111111111111:cloudwatch-alarm-template-group:cloud-watch-alarm-template-group-id
output "medialive_cloudwatch_alarm_template_group" {
  value = provider::arn::medialive_cloudwatch_alarm_template_group("cloud-watch-alarm-template-group-id")
}
