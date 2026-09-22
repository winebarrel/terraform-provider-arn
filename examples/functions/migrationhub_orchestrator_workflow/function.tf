# arn:aws:migrationhub-orchestrator:ap-northeast-1:111111111111:workflow/resource-id
output "migrationhub_orchestrator_workflow" {
  value = provider::arn::migrationhub_orchestrator_workflow("resource-id")
}
