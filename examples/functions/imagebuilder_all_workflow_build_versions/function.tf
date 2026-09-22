# arn:aws:imagebuilder:ap-northeast-1:111111111111:workflow/workflow-type/workflow-name/workflow-version/*
output "imagebuilder_all_workflow_build_versions" {
  value = provider::arn::imagebuilder_all_workflow_build_versions("workflow-type", "workflow-name", "workflow-version")
}
