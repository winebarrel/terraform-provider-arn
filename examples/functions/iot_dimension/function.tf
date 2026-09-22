# arn:aws:iot:ap-northeast-1:111111111111:dimension/dimension-name
output "iot_dimension" {
  value = provider::arn::iot_dimension("dimension-name")
}
