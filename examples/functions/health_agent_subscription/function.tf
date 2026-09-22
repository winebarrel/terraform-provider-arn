# arn:aws:health-agent:ap-northeast-1:111111111111:domain/domain-id/subscription/subscription-id
output "health_agent_subscription" {
  value = provider::arn::health_agent_subscription("domain-id", "subscription-id")
}
