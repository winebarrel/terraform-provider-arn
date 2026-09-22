# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/phone-number/phone-number-id
output "connect_legacy_phone_number" {
  value = provider::arn::connect_legacy_phone_number("instance-id", "phone-number-id")
}
