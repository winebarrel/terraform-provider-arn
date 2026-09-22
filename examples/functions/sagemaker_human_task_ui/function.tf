# arn:aws:sagemaker:ap-northeast-1:111111111111:human-task-ui/human-task-ui-name
output "sagemaker_human_task_ui" {
  value = provider::arn::sagemaker_human_task_ui("human-task-ui-name")
}
