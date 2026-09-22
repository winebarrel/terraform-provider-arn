# arn:aws:rds:ap-northeast-1:111111111111:cluster-endpoint:db-cluster-endpoint
output "rds_cluster_endpoint" {
  value = provider::arn::rds_cluster_endpoint("db-cluster-endpoint")
}
