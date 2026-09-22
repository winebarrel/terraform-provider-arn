# arn:aws:sagemaker:ap-northeast-1:111111111111:app/domain-id/user-profile-name/app-type/app-name
output "sagemaker_app" {
  value = provider::arn::sagemaker_app("domain-id", "user-profile-name", "app-type", "app-name")
}
