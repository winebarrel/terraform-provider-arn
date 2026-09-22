# arn:aws:network-firewall:ap-northeast-1:111111111111:proxy-rule-group/name
output "network_firewall_proxy_rule_group" {
  value = provider::arn::network_firewall_proxy_rule_group("name")
}
