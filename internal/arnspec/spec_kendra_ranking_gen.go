// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: kendra-ranking
// Source: https://servicereference.us-east-1.amazonaws.com/v1/kendra-ranking/kendra-ranking.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "kendra_ranking_rescore_execution_plan", Service: "kendra-ranking", Resource: "rescore-execution-plan", Template: "arn:${Partition}:kendra-ranking:${Region}:${Account}:rescore-execution-plan/${RescoreExecutionPlanId}"},
	})
}
