# arn:aws:lambda:ap-northeast-1:111111111111:function:function-name:version
output "lambda_function_version" {
  value = provider::arn::lambda_function_version("function-name", "version")
}
