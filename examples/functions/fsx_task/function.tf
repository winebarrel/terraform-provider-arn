# arn:aws:fsx:ap-northeast-1:111111111111:task/task-id
output "fsx_task" {
  value = provider::arn::fsx_task("task-id")
}
