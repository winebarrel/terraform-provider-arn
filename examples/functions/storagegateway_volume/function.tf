# arn:aws:storagegateway:ap-northeast-1:111111111111:gateway/gateway-id/volume/volume-id
output "storagegateway_volume" {
  value = provider::arn::storagegateway_volume("gateway-id", "volume-id")
}
