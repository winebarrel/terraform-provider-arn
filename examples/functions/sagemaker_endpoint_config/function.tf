# arn:aws:sagemaker:ap-northeast-1:111111111111:endpoint-config/endpoint-config-name
output "sagemaker_endpoint_config" {
  value = provider::arn::sagemaker_endpoint_config("endpoint-config-name")
}
