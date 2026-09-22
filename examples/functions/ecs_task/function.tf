# arn:aws:ecs:ap-northeast-1:111111111111:task/cluster-name/task-id
output "ecs_task" {
  value = provider::arn::ecs_task("cluster-name", "task-id")
}
