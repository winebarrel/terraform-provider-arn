# arn:aws:social-messaging:ap-northeast-1:111111111111:phone-number-id/origination-phone-number-id
output "social_messaging_phone_number_id" {
  value = provider::arn::social_messaging_phone_number_id("origination-phone-number-id")
}
