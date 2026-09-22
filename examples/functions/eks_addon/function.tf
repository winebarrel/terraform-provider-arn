# arn:aws:eks:ap-northeast-1:111111111111:addon/cluster-name/addon-name/uuid
output "eks_addon" {
  value = provider::arn::eks_addon("cluster-name", "addon-name", "uuid")
}
