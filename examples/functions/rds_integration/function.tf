# arn:aws:rds:ap-northeast-1:111111111111:integration:integration-identifier
output "rds_integration" {
  value = provider::arn::rds_integration("integration-identifier")
}
