# arn:aws:rds:ap-northeast-1:111111111111:db:db-instance-name
output "rds_db" {
  value = provider::arn::rds_db("db-instance-name")
}
