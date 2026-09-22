# arn:aws:omics:ap-northeast-1:111111111111:workflow/id/version/version-name
output "omics_workflow_version" {
  value = provider::arn::omics_workflow_version("id", "version-name")
}
