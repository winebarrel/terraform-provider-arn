# arn:aws:network-security-manager:ap-northeast-1:111111111111:template:template-id
output "network_security_manager_template" {
  value = provider::arn::network_security_manager_template("template-id")
}
