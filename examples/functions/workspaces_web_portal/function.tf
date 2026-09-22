# arn:aws:workspaces-web:ap-northeast-1:111111111111:portal/portal-id
output "workspaces_web_portal" {
  value = provider::arn::workspaces_web_portal("portal-id")
}
