# arn:aws:route53profiles:ap-northeast-1:111111111111:profile/resource-id
output "route53profiles_profile" {
  value = provider::arn::route53profiles_profile("resource-id")
}
