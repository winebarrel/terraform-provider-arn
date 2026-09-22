# arn:aws:rum:ap-northeast-1:111111111111:appmonitor/name
output "rum_app_monitor_resource" {
  value = provider::arn::rum_app_monitor_resource("name")
}
