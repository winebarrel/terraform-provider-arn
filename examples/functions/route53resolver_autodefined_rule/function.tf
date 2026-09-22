# arn:aws:route53resolver:ap-northeast-1:111111111111:autodefined-rule/resource-id
output "route53resolver_autodefined_rule" {
  value = provider::arn::route53resolver_autodefined_rule("resource-id")
}
