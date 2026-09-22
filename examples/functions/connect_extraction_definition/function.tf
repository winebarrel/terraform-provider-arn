# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/extraction-definition/extraction-definition-id
output "connect_extraction_definition" {
  value = provider::arn::connect_extraction_definition("instance-id", "extraction-definition-id")
}
