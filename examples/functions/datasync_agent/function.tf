# arn:aws:datasync:ap-northeast-1:111111111111:agent/agent-id
output "datasync_agent" {
  value = provider::arn::datasync_agent("agent-id")
}
