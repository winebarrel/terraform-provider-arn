# arn:aws:imagebuilder:ap-northeast-1:111111111111:workflow-execution/workflow-execution-id
output "imagebuilder_workflow_execution" {
  value = provider::arn::imagebuilder_workflow_execution("workflow-execution-id")
}
