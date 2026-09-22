# arn:aws:lightsail:ap-northeast-1:111111111111:RelationalDatabaseSnapshot/id
output "lightsail_relational_database_snapshot" {
  value = provider::arn::lightsail_relational_database_snapshot("id")
}
