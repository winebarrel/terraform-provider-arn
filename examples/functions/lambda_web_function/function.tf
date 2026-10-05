# arn:aws:lambda:ap-northeast-1:111111111111:web-function/function-name
output "lambda_web_function" {
  value = provider::arn::lambda_web_function("function-name")
}
