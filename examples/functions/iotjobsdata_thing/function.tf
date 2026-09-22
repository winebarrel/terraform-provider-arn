# arn:aws:iot:ap-northeast-1:111111111111:thing/thing-name
output "iotjobsdata_thing" {
  value = provider::arn::iotjobsdata_thing("thing-name")
}
