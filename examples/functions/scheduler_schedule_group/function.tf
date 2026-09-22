# arn:aws:scheduler:ap-northeast-1:111111111111:schedule-group/group-name
output "scheduler_schedule_group" {
  value = provider::arn::scheduler_schedule_group("group-name")
}
