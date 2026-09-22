# arn:aws:storagegateway:ap-northeast-1:111111111111:gateway/gateway-id
output "storagegateway_gateway" {
  value = provider::arn::storagegateway_gateway("gateway-id")
}
