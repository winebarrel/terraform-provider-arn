# arn:aws:inspector2:ap-northeast-1:111111111111:owner/owner-id/filter/filter-id
output "inspector2_filter" {
  value = provider::arn::inspector2_filter("owner-id", "filter-id")
}
