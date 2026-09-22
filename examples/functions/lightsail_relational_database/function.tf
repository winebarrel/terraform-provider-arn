# arn:aws:lightsail:ap-northeast-1:111111111111:RelationalDatabase/id
output "lightsail_relational_database" {
  value = provider::arn::lightsail_relational_database("id")
}
