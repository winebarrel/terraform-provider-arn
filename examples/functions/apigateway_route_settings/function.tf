# arn:aws:apigateway:ap-northeast-1::/apis/api-id/stages/stage-name/routesettings/route-key
output "apigateway_route_settings" {
  value = provider::arn::apigateway_route_settings("api-id", "stage-name", "route-key")
}
