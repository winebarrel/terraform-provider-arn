# arn:aws:workspaces-web:ap-northeast-1:111111111111:ipAccessSettings/ip-access-settings-id
output "workspaces_web_ip_access_settings" {
  value = provider::arn::workspaces_web_ip_access_settings("ip-access-settings-id")
}
