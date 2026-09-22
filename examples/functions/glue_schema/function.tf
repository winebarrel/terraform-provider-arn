# arn:aws:glue:ap-northeast-1:111111111111:schema/schema-name
output "glue_schema" {
  value = provider::arn::glue_schema("schema-name")
}
