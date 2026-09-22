# arn:aws:sagemaker:ap-northeast-1:111111111111:hyper-parameter-tuning-job/hyper-parameter-tuning-job-name
output "sagemaker_hyper_parameter_tuning_job" {
  value = provider::arn::sagemaker_hyper_parameter_tuning_job("hyper-parameter-tuning-job-name")
}
