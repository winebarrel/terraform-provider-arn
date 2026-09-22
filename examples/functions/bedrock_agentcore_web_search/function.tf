# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:tool/web-search.v1
output "bedrock_agentcore_web_search" {
  value = provider::arn::bedrock_agentcore_web_search()
}
