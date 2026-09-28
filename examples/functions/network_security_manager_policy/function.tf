# arn:aws:network-security-manager:ap-northeast-1:111111111111:policy:policy-id
output "network_security_manager_policy" {
  value = provider::arn::network_security_manager_policy("policy-id")
}
