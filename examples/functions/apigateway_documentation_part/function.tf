# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/documentation/parts/documentation-part-id
output "apigateway_documentation_part" {
  value = provider::arn::apigateway_documentation_part("rest-api-id", "documentation-part-id")
}
