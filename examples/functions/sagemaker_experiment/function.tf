# arn:aws:sagemaker:ap-northeast-1:111111111111:experiment/experiment-name
output "sagemaker_experiment" {
  value = provider::arn::sagemaker_experiment("experiment-name")
}
