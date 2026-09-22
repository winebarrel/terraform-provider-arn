# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/contact-flow/*
output "connect_wildcard_contact_flow" {
  value = provider::arn::connect_wildcard_contact_flow("instance-id")
}
