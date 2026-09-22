# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/transfer-destination/quick-connect-id
output "connect_quick_connect" {
  value = provider::arn::connect_quick_connect("instance-id", "quick-connect-id")
}
