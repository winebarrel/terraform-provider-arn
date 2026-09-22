# arn:aws:iotsitewise:ap-northeast-1:111111111111:gateway/gateway-id
output "iotsitewise_gateway" {
  value = provider::arn::iotsitewise_gateway("gateway-id")
}
