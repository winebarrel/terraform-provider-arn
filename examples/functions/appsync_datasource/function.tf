# arn:aws:appsync:ap-northeast-1:111111111111:apis/graph-qlapi-id/datasources/datasource-name
output "appsync_datasource" {
  value = provider::arn::appsync_datasource("graph-qlapi-id", "datasource-name")
}
