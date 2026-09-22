# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/contact/contact-id
output "connect_contact" {
  value = provider::arn::connect_contact("instance-id", "contact-id")
}
