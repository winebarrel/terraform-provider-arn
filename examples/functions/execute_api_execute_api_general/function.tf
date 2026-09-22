# arn:aws:execute-api:ap-northeast-1:111111111111:api-id/stage/method/api-specific-resource-path
output "execute_api_execute_api_general" {
  value = provider::arn::execute_api_execute_api_general("api-id", "stage", "method", "api-specific-resource-path")
}
