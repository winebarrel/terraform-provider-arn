# arn:aws:network-security-manager:ap-northeast-1:111111111111:deployment:deployment-id
output "network_security_manager_deployment" {
  value = provider::arn::network_security_manager_deployment("deployment-id")
}
