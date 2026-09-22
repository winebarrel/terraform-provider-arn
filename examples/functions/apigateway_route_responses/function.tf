# arn:aws:apigateway:ap-northeast-1::/apis/api-id/routes/route-id/routeresponses
output "apigateway_route_responses" {
  value = provider::arn::apigateway_route_responses("api-id", "route-id")
}
