# arn:aws:bedrock:ap-northeast-1:111111111111:agent/agent-id
output "bedrock_agent" {
  value = provider::arn::bedrock_agent("agent-id")
}
