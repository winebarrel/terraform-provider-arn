# arn:aws:route53profiles:ap-northeast-1:111111111111:profile-association/resource-id
output "route53profiles_profile_association" {
  value = provider::arn::route53profiles_profile_association("resource-id")
}
