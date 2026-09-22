# arn:aws:redshift:ap-northeast-1:111111111111:integration:integration-identifier
output "redshift_integration" {
  value = provider::arn::redshift_integration("integration-identifier")
}
