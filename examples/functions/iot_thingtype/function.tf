# arn:aws:iot:ap-northeast-1:111111111111:thingtype/thing-type-name
output "iot_thingtype" {
  value = provider::arn::iot_thingtype("thing-type-name")
}
