# arn:aws:ecs:ap-northeast-1:111111111111:daemon-task-definition/daemon-task-definition-family-name:daemon-task-definition-revision-number
output "ecs_daemon_task_definition" {
  value = provider::arn::ecs_daemon_task_definition("daemon-task-definition-family-name", "daemon-task-definition-revision-number")
}
