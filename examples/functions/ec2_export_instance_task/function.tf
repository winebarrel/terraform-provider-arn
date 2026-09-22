# arn:aws:ec2:ap-northeast-1:111111111111:export-instance-task/export-task-id
output "ec2_export_instance_task" {
  value = provider::arn::ec2_export_instance_task("export-task-id")
}
