# arn:aws:rds:ap-northeast-1:111111111111:db-proxy:db-proxy-id
output "rds_proxy" {
  value = provider::arn::rds_proxy("db-proxy-id")
}
