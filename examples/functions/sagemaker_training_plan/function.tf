# arn:aws:sagemaker:ap-northeast-1:111111111111:training-plan/training-plan-name
output "sagemaker_training_plan" {
  value = provider::arn::sagemaker_training_plan("training-plan-name")
}
