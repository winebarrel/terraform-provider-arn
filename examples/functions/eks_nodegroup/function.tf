# arn:aws:eks:ap-northeast-1:111111111111:nodegroup/cluster-name/nodegroup-name/uuid
output "eks_nodegroup" {
  value = provider::arn::eks_nodegroup("cluster-name", "nodegroup-name", "uuid")
}
