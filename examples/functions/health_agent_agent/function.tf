# arn:aws:health-agent:ap-northeast-1:111111111111:domain/domain-id/agent/agent-id
output "health_agent_agent" {
  value = provider::arn::health_agent_agent("domain-id", "agent-id")
}
