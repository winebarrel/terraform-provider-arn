# arn:aws:bedrock:ap-northeast-1:111111111111:custom-model-deployment/resource-id
output "bedrock_custom_model_deployment" {
  value = provider::arn::bedrock_custom_model_deployment("resource-id")
}
