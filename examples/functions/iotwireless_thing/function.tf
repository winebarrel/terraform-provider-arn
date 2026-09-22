# arn:aws:iot:ap-northeast-1:111111111111:thing/thing-name
output "iotwireless_thing" {
  value = provider::arn::iotwireless_thing("thing-name")
}
