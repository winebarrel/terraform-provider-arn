# arn:aws:eks:ap-northeast-1:111111111111:fargateprofile/cluster-name/fargate-profile-name/uuid
output "eks_fargateprofile" {
  value = provider::arn::eks_fargateprofile("cluster-name", "fargate-profile-name", "uuid")
}
