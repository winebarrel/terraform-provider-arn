# arn:aws:imagebuilder:ap-northeast-1:111111111111:workflow-step-execution/workflow-step-execution-id
output "imagebuilder_workflow_step_execution" {
  value = provider::arn::imagebuilder_workflow_step_execution("workflow-step-execution-id")
}
