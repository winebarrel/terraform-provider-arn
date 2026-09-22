# arn:aws:lambda:ap-northeast-1:111111111111:function:function-name:version/durable-execution/execution-name/execution-id
output "lambda_durable_execution" {
  value = provider::arn::lambda_durable_execution("function-name", "version", "execution-name", "execution-id")
}
