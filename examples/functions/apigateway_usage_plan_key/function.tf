# arn:aws:apigateway:ap-northeast-1::/usageplans/usage-plan-id/keys/id
output "apigateway_usage_plan_key" {
  value = provider::arn::apigateway_usage_plan_key("usage-plan-id", "id")
}
