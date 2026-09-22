# arn:aws:iotfleetwise:ap-northeast-1:111111111111:fleet/fleet-id
output "iotfleetwise_fleet" {
  value = provider::arn::iotfleetwise_fleet("fleet-id")
}
