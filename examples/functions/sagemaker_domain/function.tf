# arn:aws:sagemaker:ap-northeast-1:111111111111:domain/domain-id
output "sagemaker_domain" {
  value = provider::arn::sagemaker_domain("domain-id")
}
