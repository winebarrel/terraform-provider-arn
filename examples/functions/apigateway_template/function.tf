# arn:aws:apigateway:ap-northeast-1::/restapis/models/model-name/template
output "apigateway_template" {
  value = provider::arn::apigateway_template("model-name")
}
