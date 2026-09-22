# arn:aws:sagemaker:ap-northeast-1:111111111111:hub-content/hub-name/hub-content-type/hub-content-name
output "sagemaker_hub_content" {
  value = provider::arn::sagemaker_hub_content("hub-name", "hub-content-type", "hub-content-name")
}
