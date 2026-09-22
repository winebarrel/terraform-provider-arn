# arn:aws:securityagent:ap-northeast-1:111111111111:integration/integration-id
output "securityagent_integration" {
  value = provider::arn::securityagent_integration("integration-id")
}
