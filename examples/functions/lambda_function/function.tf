# arn:aws:lambda:ap-northeast-1:111111111111:function:function-name
output "lambda_function" {
  value = provider::arn::lambda_function("function-name")
}
