# arn:aws:finspace:ap-northeast-1:111111111111:kxEnvironment/environment-id/kxVolume/kx-volume
output "finspace_kx_volume" {
  value = provider::arn::finspace_kx_volume("environment-id", "kx-volume")
}
