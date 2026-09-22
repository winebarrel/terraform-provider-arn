# arn:aws:finspace:ap-northeast-1:111111111111:kxEnvironment/environment-id/kxCluster/kx-cluster
output "finspace_kx_cluster" {
  value = provider::arn::finspace_kx_cluster("environment-id", "kx-cluster")
}
