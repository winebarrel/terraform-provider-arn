# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/agent/user-id
output "connect_user" {
  value = provider::arn::connect_user("instance-id", "user-id")
}
