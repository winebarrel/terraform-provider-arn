# arn:aws:qbusiness:ap-northeast-1:111111111111:application/application-id/integration/integration-id
output "qbusiness_integration" {
  value = provider::arn::qbusiness_integration("application-id", "integration-id")
}
