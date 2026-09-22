# arn:aws:rds:ap-northeast-1:111111111111:cluster-snapshot:cluster-snapshot-name
output "rds_cluster_snapshot" {
  value = provider::arn::rds_cluster_snapshot("cluster-snapshot-name")
}
