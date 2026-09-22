# arn:aws:apigateway:ap-northeast-1::/apis/api-id/routes/route-id
output "apigateway_route" {
  value = provider::arn::apigateway_route("api-id", "route-id")
}
