# arn:aws:ecs:ap-northeast-1:111111111111:task-set/cluster-name/service-name/task-set-id
output "ecs_task_set" {
  value = provider::arn::ecs_task_set("cluster-name", "service-name", "task-set-id")
}
