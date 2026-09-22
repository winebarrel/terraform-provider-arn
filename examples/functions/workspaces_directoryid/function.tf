# arn:aws:workspaces:ap-northeast-1:111111111111:directory/directory-id
output "workspaces_directoryid" {
  value = provider::arn::workspaces_directoryid("directory-id")
}
