# arn:aws:bedrock:ap-northeast-1::foundation-model/resource-id
output "bedrock_foundation_model" {
  value = provider::arn::bedrock_foundation_model("resource-id")
}
