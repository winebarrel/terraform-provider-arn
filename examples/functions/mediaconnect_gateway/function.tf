# arn:aws:mediaconnect:ap-northeast-1:111111111111:gateway:gateway-id:gateway-name
output "mediaconnect_gateway" {
  value = provider::arn::mediaconnect_gateway("gateway-id", "gateway-name")
}
