# arn:aws:rds-db:ap-northeast-1:111111111111:dbuser:dbi-resource-id/db-user-name
output "rds_db_db_user" {
  value = provider::arn::rds_db_db_user("dbi-resource-id", "db-user-name")
}
