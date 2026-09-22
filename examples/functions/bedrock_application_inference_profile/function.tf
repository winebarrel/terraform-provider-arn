# arn:aws:bedrock:ap-northeast-1:111111111111:application-inference-profile/resource-id
output "bedrock_application_inference_profile" {
  value = provider::arn::bedrock_application_inference_profile("resource-id")
}
