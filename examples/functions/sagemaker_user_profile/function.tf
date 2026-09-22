# arn:aws:sagemaker:ap-northeast-1:111111111111:user-profile/domain-id/user-profile-name
output "sagemaker_user_profile" {
  value = provider::arn::sagemaker_user_profile("domain-id", "user-profile-name")
}
