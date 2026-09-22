# arn:aws:emr-containers:ap-northeast-1:111111111111:/virtualclusters/virtual-cluster-id/endpoints/endpoint-id
output "emr_containers_managed_endpoint" {
  value = provider::arn::emr_containers_managed_endpoint("virtual-cluster-id", "endpoint-id")
}
