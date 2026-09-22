# arn:aws:aidevops:ap-northeast-1:111111111111:agentspace/agent-space-id
output "aidevops_agentspace" {
  value = provider::arn::aidevops_agentspace("agent-space-id")
}
