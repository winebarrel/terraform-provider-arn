# arn:aws:appsync:ap-northeast-1:111111111111:apis/api-id
output "appsync_api" {
  value = provider::arn::appsync_api("api-id")
}
