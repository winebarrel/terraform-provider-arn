# arn:aws:lambda:ap-northeast-1:111111111111:event-source-mapping:uuid
output "lambda_event_source_mapping" {
  value = provider::arn::lambda_event_source_mapping("uuid")
}
