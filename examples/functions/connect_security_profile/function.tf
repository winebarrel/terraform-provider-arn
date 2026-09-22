# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/security-profile/security-profile-id
output "connect_security_profile" {
  value = provider::arn::connect_security_profile("instance-id", "security-profile-id")
}
