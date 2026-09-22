# arn:aws:entityresolution:ap-northeast-1:111111111111:idmappingworkflow/workflow-name
output "entityresolution_id_mapping_workflow" {
  value = provider::arn::entityresolution_id_mapping_workflow("workflow-name")
}
