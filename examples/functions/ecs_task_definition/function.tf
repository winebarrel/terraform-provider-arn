# arn:aws:ecs:ap-northeast-1:111111111111:task-definition/task-definition-family-name:task-definition-revision-number
output "ecs_task_definition" {
  value = provider::arn::ecs_task_definition("task-definition-family-name", "task-definition-revision-number")
}
