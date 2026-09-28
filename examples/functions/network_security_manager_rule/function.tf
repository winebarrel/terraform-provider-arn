# arn:aws:network-security-manager:ap-northeast-1:111111111111:rule:rule-id
output "network_security_manager_rule" {
  value = provider::arn::network_security_manager_rule("rule-id")
}
