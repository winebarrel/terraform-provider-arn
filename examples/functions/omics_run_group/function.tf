# arn:aws:omics:ap-northeast-1:111111111111:runGroup/id
output "omics_run_group" {
  value = provider::arn::omics_run_group("id")
}
