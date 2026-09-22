# arn:aws:aidevops:ap-northeast-1:111111111111:agentspace/agent-space-id/trigger/trigger-id
output "aidevops_trigger" {
  value = provider::arn::aidevops_trigger("agent-space-id", "trigger-id")
}
