# arn:aws:iotwireless:ap-northeast-1:111111111111:ServiceProfile/service-profile-id
output "iotwireless_service_profile" {
  value = provider::arn::iotwireless_service_profile("service-profile-id")
}
