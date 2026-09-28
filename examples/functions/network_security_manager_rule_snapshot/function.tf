# arn:aws:network-security-manager:ap-northeast-1:111111111111:rule:rule-id:version-number
output "network_security_manager_rule_snapshot" {
  value = provider::arn::network_security_manager_rule_snapshot("rule-id", "version-number")
}
