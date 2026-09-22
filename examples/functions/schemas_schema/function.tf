# arn:aws:schemas:ap-northeast-1:111111111111:schema/registry-name/schema-name
output "schemas_schema" {
  value = provider::arn::schemas_schema("registry-name", "schema-name")
}
