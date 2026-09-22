# arn:aws:network-firewall:ap-northeast-1:111111111111:stateless-rulegroup/name
output "network_firewall_stateless_rule_group" {
  value = provider::arn::network_firewall_stateless_rule_group("name")
}
