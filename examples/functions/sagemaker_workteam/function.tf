# arn:aws:sagemaker:ap-northeast-1:111111111111:workteam/workteam-name
output "sagemaker_workteam" {
  value = provider::arn::sagemaker_workteam("workteam-name")
}
