# arn:aws:bedrock:ap-northeast-1:111111111111:flow/flow-id
output "bedrock_flow" {
  value = provider::arn::bedrock_flow("flow-id")
}
