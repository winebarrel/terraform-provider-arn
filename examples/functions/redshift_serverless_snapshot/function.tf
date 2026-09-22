# arn:aws:redshift-serverless:ap-northeast-1:111111111111:snapshot/snapshot-id
output "redshift_serverless_snapshot" {
  value = provider::arn::redshift_serverless_snapshot("snapshot-id")
}
