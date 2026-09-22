# arn:aws:clouddirectory:ap-northeast-1:111111111111:directory/directory-id/schema/schema-name/version
output "clouddirectory_applied_schema" {
  value = provider::arn::clouddirectory_applied_schema("directory-id", "schema-name", "version")
}
