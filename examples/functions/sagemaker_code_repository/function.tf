# arn:aws:sagemaker:ap-northeast-1:111111111111:code-repository/code-repository-name
output "sagemaker_code_repository" {
  value = provider::arn::sagemaker_code_repository("code-repository-name")
}
