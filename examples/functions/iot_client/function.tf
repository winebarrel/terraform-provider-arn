# arn:aws:iot:ap-northeast-1:111111111111:client/client-id
output "iot_client" {
  value = provider::arn::iot_client("client-id")
}
