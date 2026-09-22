# arn:aws:bedrock:ap-northeast-1:111111111111:agent-alias/agent-id/agent-alias-id
output "bedrock_agent_alias" {
  value = provider::arn::bedrock_agent_alias("agent-id", "agent-alias-id")
}
