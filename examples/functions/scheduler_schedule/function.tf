# arn:aws:scheduler:ap-northeast-1:111111111111:schedule/group-name/schedule-name
output "scheduler_schedule" {
  value = provider::arn::scheduler_schedule("group-name", "schedule-name")
}
