# arn:aws:mgh:ap-northeast-1:111111111111:progressUpdateStream/stream/migrationTask/task
output "mgh_migration_task" {
  value = provider::arn::mgh_migration_task("stream", "task")
}
