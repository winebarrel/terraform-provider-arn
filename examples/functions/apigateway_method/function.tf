# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/resources/resource-id/methods/http-method-type
output "apigateway_method" {
  value = provider::arn::apigateway_method("rest-api-id", "resource-id", "http-method-type")
}
