# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/agent-state/agent-status-id
output "connect_agent_status" {
  value = provider::arn::connect_agent_status("instance-id", "agent-status-id")
}
