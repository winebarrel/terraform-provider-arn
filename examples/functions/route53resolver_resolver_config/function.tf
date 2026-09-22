# arn:aws:route53resolver:ap-northeast-1:111111111111:resolver-config/resource-id
output "route53resolver_resolver_config" {
  value = provider::arn::route53resolver_resolver_config("resource-id")
}
