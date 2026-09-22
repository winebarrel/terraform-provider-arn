# arn:aws:omics:ap-northeast-1:111111111111:run/id
output "omics_run" {
  value = provider::arn::omics_run("id")
}
