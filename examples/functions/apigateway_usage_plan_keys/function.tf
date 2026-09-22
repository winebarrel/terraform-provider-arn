# arn:aws:apigateway:ap-northeast-1::/usageplans/usage-plan-id/keys
output "apigateway_usage_plan_keys" {
  value = provider::arn::apigateway_usage_plan_keys("usage-plan-id")
}
