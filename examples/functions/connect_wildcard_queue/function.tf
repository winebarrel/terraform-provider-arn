# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/queue/*
output "connect_wildcard_queue" {
  value = provider::arn::connect_wildcard_queue("instance-id")
}
