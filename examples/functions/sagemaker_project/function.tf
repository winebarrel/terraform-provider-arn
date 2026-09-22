# arn:aws:sagemaker:ap-northeast-1:111111111111:project/project-name
output "sagemaker_project" {
  value = provider::arn::sagemaker_project("project-name")
}
