# arn:aws:network-firewall:ap-northeast-1:111111111111:firewall-policy/name
output "network_firewall_firewall_policy" {
  value = provider::arn::network_firewall_firewall_policy("name")
}
