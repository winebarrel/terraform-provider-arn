# arn:aws:rds::111111111111:global-cluster:global-cluster
output "rds_global_cluster" {
  value = provider::arn::rds_global_cluster("global-cluster")
}
