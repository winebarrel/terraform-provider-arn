# arn:aws:aidevops:ap-northeast-1:111111111111:agentspace/agent-space-id/association/association-id
output "aidevops_associations" {
  value = provider::arn::aidevops_associations("agent-space-id", "association-id")
}
