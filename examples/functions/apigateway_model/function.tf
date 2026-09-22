# arn:aws:apigateway:ap-northeast-1::/apis/api-id/models/model-id
output "apigateway_model" {
  value = provider::arn::apigateway_model("api-id", "model-id")
}
