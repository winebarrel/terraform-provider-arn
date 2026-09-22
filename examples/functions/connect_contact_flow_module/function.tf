# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/flow-module/contact-flow-module-id
output "connect_contact_flow_module" {
  value = provider::arn::connect_contact_flow_module("instance-id", "contact-flow-module-id")
}
