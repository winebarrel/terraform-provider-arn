# arn:aws:odb:ap-northeast-1:111111111111:cloud-autonomous-vm-cluster/cloud-autonomous-vm-cluster-id
output "odb_cloud_autonomous_vm_cluster" {
  value = provider::arn::odb_cloud_autonomous_vm_cluster("cloud-autonomous-vm-cluster-id")
}
