# arn:aws:omics:ap-northeast-1:111111111111:runCache/id
output "omics_run_cache" {
  value = provider::arn::omics_run_cache("id")
}
