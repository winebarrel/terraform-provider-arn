# arn:aws:route53resolver:ap-northeast-1:111111111111:resolver-query-log-config/resource-id
output "route53resolver_resolver_query_log_config" {
  value = provider::arn::route53resolver_resolver_query_log_config("resource-id")
}
