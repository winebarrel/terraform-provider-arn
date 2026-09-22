# arn:aws:lambda:ap-northeast-1:111111111111:capacity-provider:capacity-provider-name
output "lambda_capacity_provider" {
  value = provider::arn::lambda_capacity_provider("capacity-provider-name")
}
