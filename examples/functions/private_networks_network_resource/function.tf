# arn:aws:private-networks:ap-northeast-1:111111111111:network-resource/network-name/resource-id
output "private_networks_network_resource" {
  value = provider::arn::private_networks_network_resource("network-name", "resource-id")
}
