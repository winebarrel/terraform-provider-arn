# arn:aws:workspaces:ap-northeast-1:111111111111:workspaceimage/image-id
output "workspaces_workspaceimage" {
  value = provider::arn::workspaces_workspaceimage("image-id")
}
