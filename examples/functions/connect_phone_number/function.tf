# arn:aws:connect:ap-northeast-1:111111111111:phone-number/phone-number-id
output "connect_phone_number" {
  value = provider::arn::connect_phone_number("phone-number-id")
}
