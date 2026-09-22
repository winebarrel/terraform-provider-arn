# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/documentation/versions
output "apigateway_documentation_versions" {
  value = provider::arn::apigateway_documentation_versions("rest-api-id")
}
