# arn:aws:rds:ap-northeast-1:111111111111:ri:reserved-db-instance-name
output "rds_ri" {
  value = provider::arn::rds_ri("reserved-db-instance-name")
}
