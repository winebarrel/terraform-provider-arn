# arn:aws:mediaconnect:ap-northeast-1:111111111111:gateway:gateway-id:gateway-name:instance:instance-id
output "mediaconnect_gateway_instance" {
  value = provider::arn::mediaconnect_gateway_instance("gateway-id", "gateway-name", "instance-id")
}
