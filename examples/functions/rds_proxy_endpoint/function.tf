# arn:aws:rds:ap-northeast-1:111111111111:db-proxy-endpoint:db-proxy-endpoint-id
output "rds_proxy_endpoint" {
  value = provider::arn::rds_proxy_endpoint("db-proxy-endpoint-id")
}
