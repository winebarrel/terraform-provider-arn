# arn:aws:sagemaker:ap-northeast-1:111111111111:partner-app/app-id
output "sagemaker_partner_app" {
  value = provider::arn::sagemaker_partner_app("app-id")
}
