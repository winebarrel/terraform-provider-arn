# arn:aws:apigateway:ap-northeast-1::/usageplans/usage-plan-id
output "apigateway_usage_plan" {
  value = provider::arn::apigateway_usage_plan("usage-plan-id")
}
