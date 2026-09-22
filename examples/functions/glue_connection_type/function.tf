# arn:aws:glue:ap-northeast-1:111111111111:connectionType:connection-type-name
output "glue_connection_type" {
  value = provider::arn::glue_connection_type("connection-type-name")
}
