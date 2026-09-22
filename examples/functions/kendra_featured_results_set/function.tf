# arn:aws:kendra:ap-northeast-1:111111111111:index/index-id/featured-results-set/featured-results-set-id
output "kendra_featured_results_set" {
  value = provider::arn::kendra_featured_results_set("index-id", "featured-results-set-id")
}
