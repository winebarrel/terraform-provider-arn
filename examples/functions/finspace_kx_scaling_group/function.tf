# arn:aws:finspace:ap-northeast-1:111111111111:kxEnvironment/environment-id/kxScalingGroup/kx-scaling-group
output "finspace_kx_scaling_group" {
  value = provider::arn::finspace_kx_scaling_group("environment-id", "kx-scaling-group")
}
