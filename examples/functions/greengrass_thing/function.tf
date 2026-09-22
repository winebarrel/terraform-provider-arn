# arn:aws:iot:ap-northeast-1:111111111111:thing/thing-name
output "greengrass_thing" {
  value = provider::arn::greengrass_thing("thing-name")
}
