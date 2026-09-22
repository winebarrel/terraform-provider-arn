# arn:aws:apigateway:ap-northeast-1::/apis/api-id/routes/route-id/requestparameters/request-parameter-key
output "apigateway_route_request_parameter" {
  value = provider::arn::apigateway_route_request_parameter("api-id", "route-id", "request-parameter-key")
}
