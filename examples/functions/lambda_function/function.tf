# arn:aws:lambda:ap-northeast-1:111111111111:function:my-function
output "function" {
  value = provider::arn::lambda_function("my-function")
}
