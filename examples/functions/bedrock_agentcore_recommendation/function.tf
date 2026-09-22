# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:recommendation/recommendation-id
output "bedrock_agentcore_recommendation" {
  value = provider::arn::bedrock_agentcore_recommendation("recommendation-id")
}
