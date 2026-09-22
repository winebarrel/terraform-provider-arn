# arn:aws:memorydb:ap-northeast-1:111111111111:cluster/cluster-name
output "memorydb_cluster" {
  value = provider::arn::memorydb_cluster("cluster-name")
}
