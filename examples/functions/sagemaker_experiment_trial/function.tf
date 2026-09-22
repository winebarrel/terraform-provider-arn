# arn:aws:sagemaker:ap-northeast-1:111111111111:experiment-trial/trial-name
output "sagemaker_experiment_trial" {
  value = provider::arn::sagemaker_experiment_trial("trial-name")
}
