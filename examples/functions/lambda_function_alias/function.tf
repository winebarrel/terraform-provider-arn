# arn:aws:lambda:ap-northeast-1:111111111111:function:function-name:alias
output "lambda_function_alias" {
  value = provider::arn::lambda_function_alias("function-name", "alias")
}
