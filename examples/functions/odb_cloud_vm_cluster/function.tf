# arn:aws:odb:ap-northeast-1:111111111111:cloud-vm-cluster/cloud-vm-cluster-id
output "odb_cloud_vm_cluster" {
  value = provider::arn::odb_cloud_vm_cluster("cloud-vm-cluster-id")
}
