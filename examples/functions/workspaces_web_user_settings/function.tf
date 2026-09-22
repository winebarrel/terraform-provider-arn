# arn:aws:workspaces-web:ap-northeast-1:111111111111:userSettings/user-settings-id
output "workspaces_web_user_settings" {
  value = provider::arn::workspaces_web_user_settings("user-settings-id")
}
