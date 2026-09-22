# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/contact-flow/contact-flow-id
output "connect_contact_flow" {
  value = provider::arn::connect_contact_flow("instance-id", "contact-flow-id")
}
