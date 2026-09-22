# arn:aws:glue:ap-northeast-1:111111111111:devEndpoint/dev-endpoint-name
output "glue_devendpoint" {
  value = provider::arn::glue_devendpoint("dev-endpoint-name")
}
