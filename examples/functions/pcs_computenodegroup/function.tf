# arn:aws:pcs:ap-northeast-1:111111111111:cluster/cluster-identifier/computenodegroup/compute-node-group-identifier
output "pcs_computenodegroup" {
  value = provider::arn::pcs_computenodegroup("cluster-identifier", "compute-node-group-identifier")
}
