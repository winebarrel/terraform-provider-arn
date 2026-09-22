# arn:aws:lambda:ap-northeast-1:111111111111:layer:layer-name:layer-version
output "lambda_layer_version" {
  value = provider::arn::lambda_layer_version("layer-name", "layer-version")
}
