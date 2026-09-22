# arn:aws:rds:ap-northeast-1:111111111111:cluster-pg:cluster-parameter-group-name
output "rds_cluster_pg" {
  value = provider::arn::rds_cluster_pg("cluster-parameter-group-name")
}
