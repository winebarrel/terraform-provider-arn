# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/resources/resource-id
output "apigateway_resource" {
  value = provider::arn::apigateway_resource("rest-api-id", "resource-id")
}
