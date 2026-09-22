# arn:aws:iotfleetwise:ap-northeast-1:111111111111:vehicle/vehicle-id
output "iotfleetwise_vehicle" {
  value = provider::arn::iotfleetwise_vehicle("vehicle-id")
}
