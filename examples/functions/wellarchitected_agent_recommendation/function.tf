# arn:aws:wellarchitected:ap-northeast-1:111111111111:agent-recommendation/resource-id
output "wellarchitected_agent_recommendation" {
  value = provider::arn::wellarchitected_agent_recommendation("resource-id")
}
