# arn:aws:workspaces:ap-northeast-1:111111111111:workspaceapplication/work-space-application-id
output "workspaces_workspaceapplication" {
  value = provider::arn::workspaces_workspaceapplication("work-space-application-id")
}
