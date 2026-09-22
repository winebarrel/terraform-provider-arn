# arn:aws:nova-act:ap-northeast-1:111111111111:workflow-definition/workflow-definition-name/workflow-run/workflow-run-id
output "nova_act_workflow_run" {
  value = provider::arn::nova_act_workflow_run("workflow-definition-name", "workflow-run-id")
}
