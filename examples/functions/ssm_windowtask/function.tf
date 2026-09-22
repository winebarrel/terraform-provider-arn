# arn:aws:ssm:ap-northeast-1:111111111111:windowtask/window-task-id
output "ssm_windowtask" {
  value = provider::arn::ssm_windowtask("window-task-id")
}
