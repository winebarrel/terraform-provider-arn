# arn:aws:budgets::111111111111:budget/budget-name
output "budgets_budget" {
  value = provider::arn::budgets_budget("budget-name")
}
