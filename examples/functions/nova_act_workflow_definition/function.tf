# arn:aws:nova-act:ap-northeast-1:111111111111:workflow-definition/workflow-definition-name
output "nova_act_workflow_definition" {
  value = provider::arn::nova_act_workflow_definition("workflow-definition-name")
}
