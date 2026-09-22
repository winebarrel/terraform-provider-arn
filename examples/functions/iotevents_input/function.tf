# arn:aws:iotevents:ap-northeast-1:111111111111:input/input-name
output "iotevents_input" {
  value = provider::arn::iotevents_input("input-name")
}
