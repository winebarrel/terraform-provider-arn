// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: resiliencehub
// Source: https://servicereference.us-east-1.amazonaws.com/v1/resiliencehub/resiliencehub.json
// Functions: 8
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "resiliencehub_app_assessment", Service: "resiliencehub", Resource: "app-assessment", Template: "arn:${Partition}:resiliencehub:${Region}:${Account}:app-assessment/${AppAssessmentId}"},
		{Name: "resiliencehub_application", Service: "resiliencehub", Resource: "application", Template: "arn:${Partition}:resiliencehub:${Region}:${Account}:app/${AppId}"},
		{Name: "resiliencehub_policy", Service: "resiliencehub", Resource: "policy", Template: "arn:${Partition}:resiliencehub:${Region}:${Account}:policy/${PolicyId}"},
		{Name: "resiliencehub_recommendation_template", Service: "resiliencehub", Resource: "recommendation-template", Template: "arn:${Partition}:resiliencehub:${Region}:${Account}:recommendation-template/${RecommendationTemplateId}"},
		{Name: "resiliencehub_resiliency_policy", Service: "resiliencehub", Resource: "resiliency-policy", Template: "arn:${Partition}:resiliencehub:${Region}:${Account}:resiliency-policy/${ResiliencyPolicyId}"},
		{Name: "resiliencehub_service", Service: "resiliencehub", Resource: "service", Template: "arn:${Partition}:resiliencehub:${Region}:${Account}:service/${ServiceId}"},
		{Name: "resiliencehub_system", Service: "resiliencehub", Resource: "system", Template: "arn:${Partition}:resiliencehub:${Region}:${Account}:system/${SystemId}"},
		{Name: "resiliencehub_test_template", Service: "resiliencehub", Resource: "test-template", Template: "arn:${Partition}:resiliencehub:${Region}:${Account}:test-template/${TestTemplateId}"},
	})
}
