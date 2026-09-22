# arn:aws:eks:ap-northeast-1:111111111111:cluster/cluster-name
output "eks_cluster" {
  value = provider::arn::eks_cluster("cluster-name")
}
