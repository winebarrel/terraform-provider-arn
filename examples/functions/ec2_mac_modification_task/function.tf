# arn:aws:ec2:ap-northeast-1:111111111111:mac-modification-task/mac-modification-task-id
output "ec2_mac_modification_task" {
  value = provider::arn::ec2_mac_modification_task("mac-modification-task-id")
}
