# arn:aws:lambda:ap-northeast-1:111111111111:web-function/function-name/endpoint/endpoint-name
output "lambda_web_function_endpoint" {
  value = provider::arn::lambda_web_function_endpoint("function-name", "endpoint-name")
}
