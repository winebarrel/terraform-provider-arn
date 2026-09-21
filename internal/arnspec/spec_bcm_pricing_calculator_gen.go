// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: bcm-pricing-calculator
// Source: https://servicereference.us-east-1.amazonaws.com/v1/bcm-pricing-calculator/bcm-pricing-calculator.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "bcm_pricing_calculator_bill_estimate", Service: "bcm-pricing-calculator", Resource: "bill-estimate", Template: "arn:${Partition}:bcm-pricing-calculator::${Account}:bill-estimate/${BillEstimateId}"},
		{Name: "bcm_pricing_calculator_bill_scenario", Service: "bcm-pricing-calculator", Resource: "bill-scenario", Template: "arn:${Partition}:bcm-pricing-calculator::${Account}:bill-scenario/${BillScenarioId}"},
		{Name: "bcm_pricing_calculator_workload_estimate", Service: "bcm-pricing-calculator", Resource: "workload-estimate", Template: "arn:${Partition}:bcm-pricing-calculator::${Account}:workload-estimate/${WorkloadEstimateId}"},
	})
}
