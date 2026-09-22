# arn:aws:route53resolver:ap-northeast-1:111111111111:resolver-endpoint/resource-id
output "route53resolver_resolver_endpoint" {
  value = provider::arn::route53resolver_resolver_endpoint("resource-id")
}
