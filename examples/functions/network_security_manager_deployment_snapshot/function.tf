# arn:aws:network-security-manager:ap-northeast-1:111111111111:deployment:deployment-id:version-number
output "network_security_manager_deployment_snapshot" {
  value = provider::arn::network_security_manager_deployment_snapshot("deployment-id", "version-number")
}
