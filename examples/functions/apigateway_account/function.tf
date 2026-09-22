# arn:aws:apigateway:ap-northeast-1::/account
output "apigateway_account" {
  value = provider::arn::apigateway_account()
}
