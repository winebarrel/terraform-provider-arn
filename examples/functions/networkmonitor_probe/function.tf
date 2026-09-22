# arn:aws:networkmonitor:ap-northeast-1:111111111111:probe/probe-id
output "networkmonitor_probe" {
  value = provider::arn::networkmonitor_probe("probe-id")
}
