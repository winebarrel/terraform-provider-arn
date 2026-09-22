# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/integration-association/integration-association-id
output "connect_integration_association" {
  value = provider::arn::connect_integration_association("instance-id", "integration-association-id")
}
