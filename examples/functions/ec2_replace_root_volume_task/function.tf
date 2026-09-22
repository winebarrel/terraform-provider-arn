# arn:aws:ec2:ap-northeast-1:111111111111:replace-root-volume-task/replace-root-volume-task-id
output "ec2_replace_root_volume_task" {
  value = provider::arn::ec2_replace_root_volume_task("replace-root-volume-task-id")
}
