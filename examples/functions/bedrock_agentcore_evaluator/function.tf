# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:evaluator/evaluator-id
output "bedrock_agentcore_evaluator" {
  value = provider::arn::bedrock_agentcore_evaluator("evaluator-id")
}
