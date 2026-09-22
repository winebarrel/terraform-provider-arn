# arn:aws:bedrock:ap-northeast-1:111111111111:custom-model/resource-id
output "bedrock_custom_model" {
  value = provider::arn::bedrock_custom_model("resource-id")
}
