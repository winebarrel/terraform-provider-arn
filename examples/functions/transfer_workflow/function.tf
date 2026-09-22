# arn:aws:transfer:ap-northeast-1:111111111111:workflow/workflow-id
output "transfer_workflow" {
  value = provider::arn::transfer_workflow("workflow-id")
}
