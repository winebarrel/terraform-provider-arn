# arn:aws:bedrock:ap-northeast-1:111111111111:blueprint/blueprint-id
output "bedrock_blueprint" {
  value = provider::arn::bedrock_blueprint("blueprint-id")
}
