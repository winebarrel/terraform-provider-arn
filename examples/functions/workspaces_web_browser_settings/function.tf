# arn:aws:workspaces-web:ap-northeast-1:111111111111:browserSettings/browser-settings-id
output "workspaces_web_browser_settings" {
  value = provider::arn::workspaces_web_browser_settings("browser-settings-id")
}
