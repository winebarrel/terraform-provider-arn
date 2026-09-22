# arn:aws:health-agent:ap-northeast-1:111111111111:domain/domain-id/integration/integration-id
output "health_agent_integration" {
  value = provider::arn::health_agent_integration("domain-id", "integration-id")
}
