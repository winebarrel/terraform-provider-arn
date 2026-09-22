# arn:aws:dms:ap-northeast-1:111111111111:migration-project:*
output "dms_migration_project" {
  value = provider::arn::dms_migration_project()
}
