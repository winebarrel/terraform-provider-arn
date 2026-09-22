# arn:aws:storagegateway:ap-northeast-1:111111111111:gateway/gateway-id/device/vtldevice
output "storagegateway_device" {
  value = provider::arn::storagegateway_device("gateway-id", "vtldevice")
}
