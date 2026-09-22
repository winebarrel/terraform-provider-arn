# arn:aws:workspaces:ap-northeast-1:111111111111:workspaceipgroup/group-id
output "workspaces_workspaceipgroup" {
  value = provider::arn::workspaces_workspaceipgroup("group-id")
}
