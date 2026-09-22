# arn:aws:wellarchitected:ap-northeast-1:111111111111:workload/resource-id
output "wellarchitected_workload" {
  value = provider::arn::wellarchitected_workload("resource-id")
}
