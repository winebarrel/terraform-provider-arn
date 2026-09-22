# arn:aws:iot:ap-northeast-1:111111111111:thinggroup/thing-group-name
output "iot_dynamicthinggroup" {
  value = provider::arn::iot_dynamicthinggroup("thing-group-name")
}
