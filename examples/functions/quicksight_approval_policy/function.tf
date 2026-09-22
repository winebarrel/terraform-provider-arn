# arn:aws:quicksight:ap-northeast-1:111111111111:approval-policy/resource-id
output "quicksight_approval_policy" {
  value = provider::arn::quicksight_approval_policy("resource-id")
}
