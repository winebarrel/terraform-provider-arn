# arn:aws:rds:ap-northeast-1:111111111111:cluster:db-cluster-instance-name
output "rds_data_cluster" {
  value = provider::arn::rds_data_cluster("db-cluster-instance-name")
}
