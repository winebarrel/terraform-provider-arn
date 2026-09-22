# arn:aws:eks:ap-northeast-1:111111111111:capability/cluster-name/capability-type/capability-name/uuid
output "eks_capability" {
  value = provider::arn::eks_capability("cluster-name", "capability-type", "capability-name", "uuid")
}
