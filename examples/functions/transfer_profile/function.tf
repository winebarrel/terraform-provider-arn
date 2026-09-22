# arn:aws:transfer:ap-northeast-1:111111111111:profile/profile-id
output "transfer_profile" {
  value = provider::arn::transfer_profile("profile-id")
}
