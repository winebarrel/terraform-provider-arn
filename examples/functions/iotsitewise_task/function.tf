# arn:aws:iotsitewise:ap-northeast-1:111111111111:workspace/workspace-name/task/task-name
output "iotsitewise_task" {
  value = provider::arn::iotsitewise_task("workspace-name", "task-name")
}
