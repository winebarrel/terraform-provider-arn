# arn:aws:network-security-manager:ap-northeast-1:111111111111:template:template-id:version-number
output "network_security_manager_template_snapshot" {
  value = provider::arn::network_security_manager_template_snapshot("template-id", "version-number")
}
