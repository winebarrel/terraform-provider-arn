# arn:aws:clouddirectory:ap-northeast-1:111111111111:schema/development/schema-name
output "clouddirectory_development_schema" {
  value = provider::arn::clouddirectory_development_schema("schema-name")
}
