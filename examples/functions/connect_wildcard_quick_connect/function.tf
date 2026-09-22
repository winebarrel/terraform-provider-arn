# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/transfer-destination/*
output "connect_wildcard_quick_connect" {
  value = provider::arn::connect_wildcard_quick_connect("instance-id")
}
