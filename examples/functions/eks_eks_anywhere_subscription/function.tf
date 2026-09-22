# arn:aws:eks:ap-northeast-1:111111111111:eks-anywhere-subscription/uuid
output "eks_eks_anywhere_subscription" {
  value = provider::arn::eks_eks_anywhere_subscription("uuid")
}
