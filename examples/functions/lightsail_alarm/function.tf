# arn:aws:lightsail:ap-northeast-1:111111111111:Alarm/id
output "lightsail_alarm" {
  value = provider::arn::lightsail_alarm("id")
}
