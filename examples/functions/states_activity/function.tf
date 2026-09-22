# arn:aws:states:ap-northeast-1:111111111111:activity:activity-name
output "states_activity" {
  value = provider::arn::states_activity("activity-name")
}
