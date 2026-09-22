# arn:aws:apigateway:ap-northeast-1::/restapis/api-id/stages/stage-name
output "wafv2_apigateway" {
  value = provider::arn::wafv2_apigateway("api-id", "stage-name")
}
