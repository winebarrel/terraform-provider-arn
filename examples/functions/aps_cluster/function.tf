# arn:aws:eks:ap-northeast-1:111111111111:cluster/cluster-name
output "aps_cluster" {
  value = provider::arn::aps_cluster("cluster-name")
}
