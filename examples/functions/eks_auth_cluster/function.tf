# arn:aws:eks:ap-northeast-1:111111111111:cluster/cluster-name
output "eks_auth_cluster" {
  value = provider::arn::eks_auth_cluster("cluster-name")
}
