# arn:aws:appsync:ap-northeast-1:111111111111:apis/graph-qlapi-id/functions/function-id
output "appsync_function" {
  value = provider::arn::appsync_function("graph-qlapi-id", "function-id")
}
