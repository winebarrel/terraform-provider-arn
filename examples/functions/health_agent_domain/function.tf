# arn:aws:health-agent:ap-northeast-1:111111111111:domain/domain-id
output "health_agent_domain" {
  value = provider::arn::health_agent_domain("domain-id")
}
