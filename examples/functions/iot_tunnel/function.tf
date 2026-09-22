# arn:aws:iot:ap-northeast-1:111111111111:tunnel/tunnel-id
output "iot_tunnel" {
  value = provider::arn::iot_tunnel("tunnel-id")
}
