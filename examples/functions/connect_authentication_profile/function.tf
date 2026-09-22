# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/authentication-profile/authentication-profile-id
output "connect_authentication_profile" {
  value = provider::arn::connect_authentication_profile("instance-id", "authentication-profile-id")
}
