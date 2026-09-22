# arn:aws:cloudwatch:ap-northeast-1:111111111111:access-profile/profile-id
output "cloudwatch_access_profile" {
  value = provider::arn::cloudwatch_access_profile("profile-id")
}
