# arn:aws:network-firewall:ap-northeast-1:111111111111:proxy/name
output "network_firewall_proxy" {
  value = provider::arn::network_firewall_proxy("name")
}
