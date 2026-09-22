# arn:aws:rds:ap-northeast-1:111111111111:cluster-auto-backup:db-cluster-automated-backup-id
output "rds_cluster_auto_backup" {
  value = provider::arn::rds_cluster_auto_backup("db-cluster-automated-backup-id")
}
