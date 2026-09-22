# arn:aws:omics:ap-northeast-1:111111111111:task/id
output "omics_task_resource" {
  value = provider::arn::omics_task_resource("id")
}
