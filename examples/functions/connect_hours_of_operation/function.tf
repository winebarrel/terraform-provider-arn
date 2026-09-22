# arn:aws:connect:ap-northeast-1:111111111111:instance/instance-id/operating-hours/hours-of-operation-id
output "connect_hours_of_operation" {
  value = provider::arn::connect_hours_of_operation("instance-id", "hours-of-operation-id")
}
