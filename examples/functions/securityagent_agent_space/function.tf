# arn:aws:securityagent:ap-northeast-1:111111111111:agent-space/agent-id
output "securityagent_agent_space" {
  value = provider::arn::securityagent_agent_space("agent-id")
}
