# arn:aws:network-security-manager:ap-northeast-1:111111111111:policy:policy-id:version-number
output "network_security_manager_policy_snapshot" {
  value = provider::arn::network_security_manager_policy_snapshot("policy-id", "version-number")
}
