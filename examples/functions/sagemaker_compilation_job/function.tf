# arn:aws:sagemaker:ap-northeast-1:111111111111:compilation-job/compilation-job-name
output "sagemaker_compilation_job" {
  value = provider::arn::sagemaker_compilation_job("compilation-job-name")
}
