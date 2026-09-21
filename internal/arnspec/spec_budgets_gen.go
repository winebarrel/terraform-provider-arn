// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: budgets
// Source: https://servicereference.us-east-1.amazonaws.com/v1/budgets/budgets.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "budgets_budget", Service: "budgets", Resource: "budget", Template: "arn:${Partition}:budgets::${Account}:budget/${BudgetName}"},
		{Name: "budgets_budget_action", Service: "budgets", Resource: "budgetAction", Template: "arn:${Partition}:budgets::${Account}:budget/${BudgetName}/action/${ActionId}"},
	})
}
