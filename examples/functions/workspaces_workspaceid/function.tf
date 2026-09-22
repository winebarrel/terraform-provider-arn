# arn:aws:workspaces:ap-northeast-1:111111111111:workspace/workspace-id
output "workspaces_workspaceid" {
  value = provider::arn::workspaces_workspaceid("workspace-id")
}
