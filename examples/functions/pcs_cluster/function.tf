# arn:aws:pcs:ap-northeast-1:111111111111:cluster/cluster-identifier
output "pcs_cluster" {
  value = provider::arn::pcs_cluster("cluster-identifier")
}
