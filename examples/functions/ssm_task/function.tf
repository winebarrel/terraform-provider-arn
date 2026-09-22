# arn:aws:ecs:ap-northeast-1:111111111111:task/task-id
output "ssm_task" {
  value = provider::arn::ssm_task("task-id")
}
