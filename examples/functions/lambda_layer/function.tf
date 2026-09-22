# arn:aws:lambda:ap-northeast-1:111111111111:layer:layer-name
output "lambda_layer" {
  value = provider::arn::lambda_layer("layer-name")
}
