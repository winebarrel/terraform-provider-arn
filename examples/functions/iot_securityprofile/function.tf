# arn:aws:iot:ap-northeast-1:111111111111:securityprofile/security-profile-name
output "iot_securityprofile" {
  value = provider::arn::iot_securityprofile("security-profile-name")
}
