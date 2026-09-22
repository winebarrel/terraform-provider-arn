# arn:aws:emr-containers:ap-northeast-1:111111111111:/virtualclusters/virtual-cluster-id
output "emr_containers_virtual_cluster" {
  value = provider::arn::emr_containers_virtual_cluster("virtual-cluster-id")
}
