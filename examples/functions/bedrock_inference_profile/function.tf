# arn:aws:bedrock:ap-northeast-1:111111111111:inference-profile/resource-id
output "bedrock_inference_profile" {
  value = provider::arn::bedrock_inference_profile("resource-id")
}
