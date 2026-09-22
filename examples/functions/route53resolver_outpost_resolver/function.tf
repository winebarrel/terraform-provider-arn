# arn:aws:route53resolver:ap-northeast-1:111111111111:outpost-resolver/resource-id
output "route53resolver_outpost_resolver" {
  value = provider::arn::route53resolver_outpost_resolver("resource-id")
}
