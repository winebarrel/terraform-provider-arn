# arn:aws:backup-search:ap-northeast-1:111111111111:search-export-job/resource-id
output "backup_search_search_export_job" {
  value = provider::arn::backup_search_search_export_job("resource-id")
}
