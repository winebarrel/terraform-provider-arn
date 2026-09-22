# arn:aws:migrationhub-orchestrator:ap-northeast-1:111111111111:template/resource-id
output "migrationhub_orchestrator_template" {
  value = provider::arn::migrationhub_orchestrator_template("resource-id")
}
