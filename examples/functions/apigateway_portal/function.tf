# arn:aws:apigateway:ap-northeast-1:111111111111:/portals/portal-id
output "apigateway_portal" {
  value = provider::arn::apigateway_portal("portal-id")
}
