# arn:aws:glue:ap-northeast-1:111111111111:workflow/workflow-name
output "glue_workflow" {
  value = provider::arn::glue_workflow("workflow-name")
}
