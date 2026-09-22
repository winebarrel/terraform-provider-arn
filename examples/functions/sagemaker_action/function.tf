# arn:aws:sagemaker:ap-northeast-1:111111111111:action/action-name
output "sagemaker_action" {
  value = provider::arn::sagemaker_action("action-name")
}
