# arn:aws:apigateway:ap-northeast-1::/apis/api-id/exports/specification
output "apigateway_exported_api" {
  value = provider::arn::apigateway_exported_api("api-id", "specification")
}
