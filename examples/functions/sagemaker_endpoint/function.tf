# arn:aws:sagemaker:ap-northeast-1:111111111111:endpoint/endpoint-name
output "sagemaker_endpoint" {
  value = provider::arn::sagemaker_endpoint("endpoint-name")
}
