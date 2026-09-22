# arn:aws:entityresolution:ap-northeast-1:111111111111:matchingworkflow/workflow-name
output "entityresolution_matching_workflow" {
  value = provider::arn::entityresolution_matching_workflow("workflow-name")
}
