# arn:aws:medialive:ap-northeast-1:111111111111:cloudwatch-alarm-template:cloud-watch-alarm-template-id
output "medialive_cloudwatch_alarm_template" {
  value = provider::arn::medialive_cloudwatch_alarm_template("cloud-watch-alarm-template-id")
}
