# arn:aws:workspaces-web:ap-northeast-1:111111111111:networkSettings/network-settings-id
output "workspaces_web_network_settings" {
  value = provider::arn::workspaces_web_network_settings("network-settings-id")
}
