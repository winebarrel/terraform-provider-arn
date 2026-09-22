# arn:aws:private-networks:ap-northeast-1:111111111111:network/network-name
output "private_networks_network" {
  value = provider::arn::private_networks_network("network-name")
}
