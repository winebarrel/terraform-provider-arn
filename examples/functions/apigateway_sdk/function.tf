# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/stages/stage-name/sdks/sdk-type
output "apigateway_sdk" {
  value = provider::arn::apigateway_sdk("rest-api-id", "stage-name", "sdk-type")
}
