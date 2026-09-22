# arn:aws:apigateway:ap-northeast-1:111111111111:/domainnames/domain-name/routingrules/routing-rule-id
output "apigateway_routing_rule" {
  value = provider::arn::apigateway_routing_rule("domain-name", "routing-rule-id")
}
