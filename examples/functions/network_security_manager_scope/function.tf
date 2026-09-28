# arn:aws:network-security-manager:ap-northeast-1:111111111111:scope:scope-id
output "network_security_manager_scope" {
  value = provider::arn::network_security_manager_scope("scope-id")
}
