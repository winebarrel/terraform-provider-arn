# arn:aws:sagemaker:ap-northeast-1:111111111111:space/domain-id/space-name
output "sagemaker_space" {
  value = provider::arn::sagemaker_space("domain-id", "space-name")
}
