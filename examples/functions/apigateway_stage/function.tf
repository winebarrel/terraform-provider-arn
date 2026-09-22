# arn:aws:apigateway:ap-northeast-1::/apis/api-id/stages/stage-name
output "apigateway_stage" {
  value = provider::arn::apigateway_stage("api-id", "stage-name")
}
