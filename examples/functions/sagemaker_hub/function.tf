# arn:aws:sagemaker:ap-northeast-1:111111111111:hub/hub-name
output "sagemaker_hub" {
  value = provider::arn::sagemaker_hub("hub-name")
}
