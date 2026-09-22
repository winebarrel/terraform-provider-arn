# arn:aws:apigateway:ap-northeast-1::/restapis/rest-api-id/stages/stage-name
output "apigateway_stage_2" {
  value = provider::arn::apigateway_stage_2("rest-api-id", "stage-name")
}
