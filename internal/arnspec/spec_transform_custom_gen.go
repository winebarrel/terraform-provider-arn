// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: transform-custom
// Source: https://servicereference.us-east-1.amazonaws.com/v1/transform-custom/transform-custom.json
// Functions: 8
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "transform_custom_analysis", Service: "transform-custom", Resource: "analysis", Template: "arn:${Partition}:transform-custom:${Region}:${Account}:analysis/${AnalysisId}"},
		{Name: "transform_custom_campaign", Service: "transform-custom", Resource: "campaign", Template: "arn:${Partition}:transform-custom:${Region}:${Account}:campaign/${Name}"},
		{Name: "transform_custom_finding", Service: "transform-custom", Resource: "finding", Template: "arn:${Partition}:transform-custom:${Region}:${Account}:finding/${FindingId}"},
		{Name: "transform_custom_knowledge_item", Service: "transform-custom", Resource: "knowledge-item", Template: "arn:${Partition}:transform-custom:${Region}:${Account}:package/${TransformationPackageName}/knowledge-item/${Id}"},
		{Name: "transform_custom_package", Service: "transform-custom", Resource: "package", Template: "arn:${Partition}:transform-custom:${Region}:${Account}:package/${Name}"},
		{Name: "transform_custom_remediation", Service: "transform-custom", Resource: "remediation", Template: "arn:${Partition}:transform-custom:${Region}:${Account}:remediation/${RemediationId}"},
		{Name: "transform_custom_repository", Service: "transform-custom", Resource: "repository", Template: "arn:${Partition}:transform-custom:${Region}:${Account}:repository/${RepositoryId}"},
		{Name: "transform_custom_source", Service: "transform-custom", Resource: "source", Template: "arn:${Partition}:transform-custom:${Region}:${Account}:source/${Name}"},
	})
}
