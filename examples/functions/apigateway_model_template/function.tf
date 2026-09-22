# arn:aws:apigateway:ap-northeast-1::/apis/api-id/models/model-id/template
output "apigateway_model_template" {
  value = provider::arn::apigateway_model_template("api-id", "model-id")
}
