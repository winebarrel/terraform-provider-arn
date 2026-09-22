# arn:aws:workspaces:ap-northeast-1:111111111111:workspacespool/pool-id
output "workspaces_workspacespool" {
  value = provider::arn::workspaces_workspacespool("pool-id")
}
