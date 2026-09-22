# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:online-evaluation-config/online-evaluation-config-id
output "bedrock_agentcore_online_evaluation_config" {
  value = provider::arn::bedrock_agentcore_online_evaluation_config("online-evaluation-config-id")
}
