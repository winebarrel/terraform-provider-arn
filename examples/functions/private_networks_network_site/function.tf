# arn:aws:private-networks:ap-northeast-1:111111111111:network-site/network-name/network-site-name
output "private_networks_network_site" {
  value = provider::arn::private_networks_network_site("network-name", "network-site-name")
}
