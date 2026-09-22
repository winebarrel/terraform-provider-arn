# arn:aws:ec2:ap-northeast-1:111111111111:export-image-task/export-image-task-id
output "ec2_export_image_task" {
  value = provider::arn::ec2_export_image_task("export-image-task-id")
}
