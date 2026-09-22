# arn:aws:ec2:ap-northeast-1:111111111111:import-snapshot-task/import-snapshot-task-id
output "ec2_import_snapshot_task" {
  value = provider::arn::ec2_import_snapshot_task("import-snapshot-task-id")
}
