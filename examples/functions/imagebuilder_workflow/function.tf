# arn:aws:imagebuilder:ap-northeast-1:111111111111:workflow/workflow-type/workflow-name/workflow-version/workflow-build-version
output "imagebuilder_workflow" {
  value = provider::arn::imagebuilder_workflow("workflow-type", "workflow-name", "workflow-version", "workflow-build-version")
}
