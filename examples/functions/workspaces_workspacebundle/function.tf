# arn:aws:workspaces:ap-northeast-1:111111111111:workspacebundle/bundle-id
output "workspaces_workspacebundle" {
  value = provider::arn::workspaces_workspacebundle("bundle-id")
}
