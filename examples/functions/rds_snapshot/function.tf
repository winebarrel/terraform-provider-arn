# arn:aws:rds:ap-northeast-1:111111111111:snapshot:snapshot-name
output "rds_snapshot" {
  value = provider::arn::rds_snapshot("snapshot-name")
}
