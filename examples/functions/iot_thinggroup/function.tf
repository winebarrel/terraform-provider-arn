# arn:aws:iot:ap-northeast-1:111111111111:thinggroup/thing-group-name
output "iot_thinggroup" {
  value = provider::arn::iot_thinggroup("thing-group-name")
}
