# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/agent-state/*
output "connect_wildcard_agent_status" {
  value = provider::arn::connect_wildcard_agent_status("instance-id")
}
