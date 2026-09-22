# arn:aws:route53resolver:ap-northeast-1:111111111111:firewall-config/resource-id
output "route53resolver_firewall_config" {
  value = provider::arn::route53resolver_firewall_config("resource-id")
}
