# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/documentation/versions/documentation-version-id
output "apigateway_documentation_version" {
  value = provider::arn::apigateway_documentation_version("rest-api-id", "documentation-version-id")
}
