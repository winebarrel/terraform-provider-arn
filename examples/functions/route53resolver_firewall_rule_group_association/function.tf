# arn:aws:route53resolver:ap-northeast-1:111111111111:firewall-rule-group-association/resource-id
output "route53resolver_firewall_rule_group_association" {
  value = provider::arn::route53resolver_firewall_rule_group_association("resource-id")
}
