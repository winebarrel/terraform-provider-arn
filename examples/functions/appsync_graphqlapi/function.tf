# arn:aws:appsync:ap-northeast-1:111111111111:apis/graph-qlapi-id
output "appsync_graphqlapi" {
  value = provider::arn::appsync_graphqlapi("graph-qlapi-id")
}
