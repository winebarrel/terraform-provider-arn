# arn:aws:groundstation:ap-northeast-1:111111111111:agent/agent-id
output "groundstation_agent" {
  value = provider::arn::groundstation_agent("agent-id")
}
