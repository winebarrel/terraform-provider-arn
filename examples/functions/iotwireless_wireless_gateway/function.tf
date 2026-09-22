# arn:aws:iotwireless:ap-northeast-1:111111111111:WirelessGateway/wireless-gateway-id
output "iotwireless_wireless_gateway" {
  value = provider::arn::iotwireless_wireless_gateway("wireless-gateway-id")
}
