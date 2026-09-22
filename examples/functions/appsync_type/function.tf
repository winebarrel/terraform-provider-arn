# arn:aws:appsync:ap-northeast-1:111111111111:apis/graph-qlapi-id/types/type-name
output "appsync_type" {
  value = provider::arn::appsync_type("graph-qlapi-id", "type-name")
}
