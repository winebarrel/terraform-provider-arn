# arn:aws:network-firewall:ap-northeast-1:111111111111:firewall/name
output "network_firewall_firewall" {
  value = provider::arn::network_firewall_firewall("name")
}
