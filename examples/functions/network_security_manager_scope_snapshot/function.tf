# arn:aws:network-security-manager:ap-northeast-1:111111111111:scope:scope-id:version-number
output "network_security_manager_scope_snapshot" {
  value = provider::arn::network_security_manager_scope_snapshot("scope-id", "version-number")
}
