# arn:aws:iot:ap-northeast-1:111111111111:ruledestination/destination-type/uuid
output "iot_destination" {
  value = provider::arn::iot_destination("destination-type", "uuid")
}
