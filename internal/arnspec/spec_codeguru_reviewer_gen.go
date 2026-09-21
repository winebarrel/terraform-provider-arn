// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: codeguru-reviewer
// Source: https://servicereference.us-east-1.amazonaws.com/v1/codeguru-reviewer/codeguru-reviewer.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "codeguru_reviewer_association", Service: "codeguru-reviewer", Resource: "association", Template: "arn:${Partition}:codeguru-reviewer:${Region}:${Account}:association:${ResourceId}"},
		{Name: "codeguru_reviewer_codereview", Service: "codeguru-reviewer", Resource: "codereview", Template: "arn:${Partition}:codeguru-reviewer:${Region}:${Account}:association:${ResourceId}:codereview:${CodeReviewId}"},
	})
}
