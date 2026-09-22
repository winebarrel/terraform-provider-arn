# arn:aws:bedrock:ap-northeast-1:111111111111:blueprint-optimization-invocation/resource-id
output "bedrock_blueprint_optimization_invocation" {
  value = provider::arn::bedrock_blueprint_optimization_invocation("resource-id")
}
