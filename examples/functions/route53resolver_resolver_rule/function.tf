# arn:aws:route53resolver:ap-northeast-1:111111111111:resolver-rule/resource-id
output "route53resolver_resolver_rule" {
  value = provider::arn::route53resolver_resolver_rule("resource-id")
}
