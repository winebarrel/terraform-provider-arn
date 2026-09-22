# arn:aws:iot:ap-northeast-1:111111111111:thing/thing-name
output "iot_thing" {
  value = provider::arn::iot_thing("thing-name")
}
