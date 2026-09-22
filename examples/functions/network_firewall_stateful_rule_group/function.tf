# arn:aws:network-firewall:ap-northeast-1:111111111111:stateful-rulegroup/name
output "network_firewall_stateful_rule_group" {
  value = provider::arn::network_firewall_stateful_rule_group("name")
}
