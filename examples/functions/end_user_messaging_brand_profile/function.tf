# arn:aws:end-user-messaging:ap-northeast-1:111111111111:brand-profile/resource-id
output "end_user_messaging_brand_profile" {
  value = provider::arn::end_user_messaging_brand_profile("resource-id")
}
