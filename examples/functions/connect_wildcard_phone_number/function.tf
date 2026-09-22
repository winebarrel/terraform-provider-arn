# arn:aws:connect:ap-northeast-1:111111111111:phone-number/*
output "connect_wildcard_phone_number" {
  value = provider::arn::connect_wildcard_phone_number()
}
