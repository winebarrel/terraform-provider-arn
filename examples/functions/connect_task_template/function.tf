# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/task-template/task-template-id
output "connect_task_template" {
  value = provider::arn::connect_task_template("instance-id", "task-template-id")
}
