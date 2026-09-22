# arn:aws:ec2:ap-northeast-1:111111111111:import-image-task/import-image-task-id
output "ec2_import_image_task" {
  value = provider::arn::ec2_import_image_task("import-image-task-id")
}
