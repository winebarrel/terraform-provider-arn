# arn:aws:datasync:ap-northeast-1:111111111111:task/task-id
output "datasync_task" {
  value = provider::arn::datasync_task("task-id")
}
