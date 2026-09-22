# arn:aws:glue:ap-northeast-1:111111111111:integration:integration-id
output "glue_integration" {
  value = provider::arn::glue_integration("integration-id")
}
