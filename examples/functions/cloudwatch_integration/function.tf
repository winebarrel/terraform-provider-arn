# arn:aws:cloudwatch:ap-northeast-1:111111111111:integration/integration-id
output "cloudwatch_integration" {
  value = provider::arn::cloudwatch_integration("integration-id")
}
