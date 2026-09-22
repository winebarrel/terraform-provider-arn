# arn:aws:apigateway:ap-northeast-1::/apis/api-id/routes
output "apigateway_routes" {
  value = provider::arn::apigateway_routes("api-id")
}
