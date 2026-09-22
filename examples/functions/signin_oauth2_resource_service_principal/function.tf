# arn:aws:signin:ap-northeast-1:111111111111:service-principal/service-principal-name
output "signin_oauth2_resource_service_principal" {
  value = provider::arn::signin_oauth2_resource_service_principal("service-principal-name")
}
