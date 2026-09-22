# arn:aws:bedrock:ap-northeast-1:111111111111:marketplace/model-endpoint/all-access
output "bedrock_bedrock_marketplace_model_endpoint" {
  value = provider::arn::bedrock_bedrock_marketplace_model_endpoint()
}
