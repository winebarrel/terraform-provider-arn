# arn:aws:iq:ap-northeast-1::buyer/buyer-id
output "iq_buyer" {
  value = provider::arn::iq_buyer("buyer-id")
}
