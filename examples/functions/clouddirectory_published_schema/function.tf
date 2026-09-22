# arn:aws:clouddirectory:ap-northeast-1:111111111111:schema/published/schema-name/version
output "clouddirectory_published_schema" {
  value = provider::arn::clouddirectory_published_schema("schema-name", "version")
}
