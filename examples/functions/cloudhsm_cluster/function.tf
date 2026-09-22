# arn:aws:cloudhsm:ap-northeast-1:111111111111:cluster/cloud-hsm-cluster-instance-name
output "cloudhsm_cluster" {
  value = provider::arn::cloudhsm_cluster("cloud-hsm-cluster-instance-name")
}
