// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: billingconductor
// Source: https://servicereference.us-east-1.amazonaws.com/v1/billingconductor/billingconductor.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "billingconductor_billinggroup", Service: "billingconductor", Resource: "billinggroup", Template: "arn:${Partition}:billingconductor::${Account}:billinggroup/${BillingGroupId}"},
		{Name: "billingconductor_customlineitem", Service: "billingconductor", Resource: "customlineitem", Template: "arn:${Partition}:billingconductor::${Account}:customlineitem/${CustomLineItemId}"},
		{Name: "billingconductor_pricingplan", Service: "billingconductor", Resource: "pricingplan", Template: "arn:${Partition}:billingconductor::${Account}:pricingplan/${PricingPlanId}"},
		{Name: "billingconductor_pricingrule", Service: "billingconductor", Resource: "pricingrule", Template: "arn:${Partition}:billingconductor::${Account}:pricingrule/${PricingRuleId}"},
	})
}
