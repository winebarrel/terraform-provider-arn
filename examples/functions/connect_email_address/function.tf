# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/email-address/email-address-id
output "connect_email_address" {
  value = provider::arn::connect_email_address("instance-id", "email-address-id")
}
