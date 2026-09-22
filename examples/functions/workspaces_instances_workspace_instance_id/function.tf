# arn:aws:workspaces-instances:ap-northeast-1:111111111111:workspaceinstance/workspace-instance-id
output "workspaces_instances_workspace_instance_id" {
  value = provider::arn::workspaces_instances_workspace_instance_id("workspace-instance-id")
}
