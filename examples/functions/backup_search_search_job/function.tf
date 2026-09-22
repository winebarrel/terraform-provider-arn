# arn:aws:backup-search:ap-northeast-1:111111111111:search-job/resource-id
output "backup_search_search_job" {
  value = provider::arn::backup_search_search_job("resource-id")
}
