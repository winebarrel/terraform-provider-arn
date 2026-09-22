# arn:aws:iotwireless:ap-northeast-1:111111111111:ImportTask/import-task-id
output "iotwireless_import_task" {
  value = provider::arn::iotwireless_import_task("import-task-id")
}
