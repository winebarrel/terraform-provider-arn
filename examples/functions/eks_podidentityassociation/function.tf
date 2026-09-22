# arn:aws:eks:ap-northeast-1:111111111111:podidentityassociation/cluster-name/uuid
output "eks_podidentityassociation" {
  value = provider::arn::eks_podidentityassociation("cluster-name", "uuid")
}
