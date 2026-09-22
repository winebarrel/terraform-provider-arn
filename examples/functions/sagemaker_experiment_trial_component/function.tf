# arn:aws:sagemaker:ap-northeast-1:111111111111:experiment-trial-component/trial-component-name
output "sagemaker_experiment_trial_component" {
  value = provider::arn::sagemaker_experiment_trial_component("trial-component-name")
}
