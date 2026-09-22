# arn:aws:apigateway:ap-northeast-1::/apis/api-id/routes/route-id/routeresponses/route-response-id
output "apigateway_route_response" {
  value = provider::arn::apigateway_route_response("api-id", "route-id", "route-response-id")
}
