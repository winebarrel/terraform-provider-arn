# arn:aws:pcs:ap-northeast-1:111111111111:cluster/cluster-identifier/queue/queue-identifier
output "pcs_queue" {
  value = provider::arn::pcs_queue("cluster-identifier", "queue-identifier")
}
