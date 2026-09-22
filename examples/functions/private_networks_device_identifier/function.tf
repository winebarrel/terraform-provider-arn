# arn:aws:private-networks:ap-northeast-1:111111111111:device-identifier/network-name/device-id
output "private_networks_device_identifier" {
  value = provider::arn::private_networks_device_identifier("network-name", "device-id")
}
