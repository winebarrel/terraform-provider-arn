# arn:aws:apigateway:ap-northeast-1::/usageplans
output "apigateway_usage_plans" {
  value = provider::arn::apigateway_usage_plans()
}
