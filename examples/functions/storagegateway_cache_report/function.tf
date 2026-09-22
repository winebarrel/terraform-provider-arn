# arn:aws:storagegateway:ap-northeast-1:111111111111:share/share-id/cache-report/cache-report-id
output "storagegateway_cache_report" {
  value = provider::arn::storagegateway_cache_report("share-id", "cache-report-id")
}
