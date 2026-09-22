# arn:aws:omics:ap-northeast-1:111111111111:workflow/id
output "omics_workflow" {
  value = provider::arn::omics_workflow("id")
}
