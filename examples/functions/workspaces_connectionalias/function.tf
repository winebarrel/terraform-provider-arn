# arn:aws:workspaces:ap-northeast-1:111111111111:connectionalias/connection-alias-id
output "workspaces_connectionalias" {
  value = provider::arn::workspaces_connectionalias("connection-alias-id")
}
