# arn:aws:batch:ap-northeast-1:111111111111:scheduling-policy/scheduling-policy-name
output "batch_scheduling_policy" {
  value = provider::arn::batch_scheduling_policy("scheduling-policy-name")
}
