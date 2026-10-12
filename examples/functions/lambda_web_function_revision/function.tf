# arn:aws:lambda:ap-northeast-1:111111111111:web-function/function-name/revision/revision-id
output "lambda_web_function_revision" {
  value = provider::arn::lambda_web_function_revision("function-name", "revision-id")
}
