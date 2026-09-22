# arn:aws:odb:ap-northeast-1:111111111111:autonomous-database/autonomous-database-id
output "odb_autonomous_database" {
  value = provider::arn::odb_autonomous_database("autonomous-database-id")
}
