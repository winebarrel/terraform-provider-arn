# arn:aws:odb:ap-northeast-1:111111111111:exadb-vm-cluster/exadb-vm-cluster-id
output "odb_exadb_vm_cluster" {
  value = provider::arn::odb_exadb_vm_cluster("exadb-vm-cluster-id")
}
