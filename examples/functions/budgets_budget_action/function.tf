# arn:aws:budgets::111111111111:budget/budget-name/action/action-id
output "budgets_budget_action" {
  value = provider::arn::budgets_budget_action("budget-name", "action-id")
}
