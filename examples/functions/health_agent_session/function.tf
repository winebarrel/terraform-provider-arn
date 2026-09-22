# arn:aws:health-agent:ap-northeast-1:111111111111:domain/domain-id/session/session-id
output "health_agent_session" {
  value = provider::arn::health_agent_session("domain-id", "session-id")
}
