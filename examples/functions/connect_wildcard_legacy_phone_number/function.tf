# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/phone-number/*
output "connect_wildcard_legacy_phone_number" {
  value = provider::arn::connect_wildcard_legacy_phone_number("instance-id")
}
