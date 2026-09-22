# arn:aws:apigateway:ap-northeast-1::/apis/api-id/authorizers/authorizer-id
output "apigateway_authorizer" {
  value = provider::arn::apigateway_authorizer("api-id", "authorizer-id")
}
