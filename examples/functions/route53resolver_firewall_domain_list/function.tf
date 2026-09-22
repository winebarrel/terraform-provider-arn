# arn:aws:route53resolver:ap-northeast-1:111111111111:firewall-domain-list/resource-id
output "route53resolver_firewall_domain_list" {
  value = provider::arn::route53resolver_firewall_domain_list("resource-id")
}
