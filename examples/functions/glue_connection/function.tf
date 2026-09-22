# arn:aws:glue:ap-northeast-1:111111111111:connection/connection-name
output "glue_connection" {
  value = provider::arn::glue_connection("connection-name")
}
