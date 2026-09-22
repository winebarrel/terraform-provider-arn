# arn:aws:elasticmapreduce:ap-northeast-1:111111111111:cluster/cluster-id
output "elasticmapreduce_cluster" {
  value = provider::arn::elasticmapreduce_cluster("cluster-id")
}
