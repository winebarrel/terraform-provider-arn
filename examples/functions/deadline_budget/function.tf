# arn:aws:deadline:ap-northeast-1:111111111111:farm/farm-id/budget/budget-id
output "deadline_budget" {
  value = provider::arn::deadline_budget("farm-id", "budget-id")
}
