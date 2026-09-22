# arn:aws:datasync:ap-northeast-1:111111111111:task/task-id/execution/execution-id
output "datasync_taskexecution" {
  value = provider::arn::datasync_taskexecution("task-id", "execution-id")
}
