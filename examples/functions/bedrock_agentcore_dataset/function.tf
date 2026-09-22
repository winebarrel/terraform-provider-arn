# arn:aws:bedrock-agentcore:ap-northeast-1:111111111111:dataset/dataset-id
output "bedrock_agentcore_dataset" {
  value = provider::arn::bedrock_agentcore_dataset("dataset-id")
}
