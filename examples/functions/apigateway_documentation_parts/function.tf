# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/documentation/parts
output "apigateway_documentation_parts" {
  value = provider::arn::apigateway_documentation_parts("rest-api-id")
}
