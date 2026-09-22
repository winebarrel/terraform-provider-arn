# arn:aws:sagemaker:ap-northeast-1:111111111111:algorithm/algorithm-name
output "sagemaker_algorithm" {
  value = provider::arn::sagemaker_algorithm("algorithm-name")
}
