// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: evidently
// Source: https://servicereference.us-east-1.amazonaws.com/v1/evidently/evidently.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "evidently_experiment", Service: "evidently", Resource: "Experiment", Template: "arn:${Partition}:evidently:${Region}:${Account}:project/${ProjectName}/experiment/${ExperimentName}"},
		{Name: "evidently_feature", Service: "evidently", Resource: "Feature", Template: "arn:${Partition}:evidently:${Region}:${Account}:project/${ProjectName}/feature/${FeatureName}"},
		{Name: "evidently_launch", Service: "evidently", Resource: "Launch", Template: "arn:${Partition}:evidently:${Region}:${Account}:project/${ProjectName}/launch/${LaunchName}"},
		{Name: "evidently_project", Service: "evidently", Resource: "Project", Template: "arn:${Partition}:evidently:${Region}:${Account}:project/${ProjectName}"},
		{Name: "evidently_segment", Service: "evidently", Resource: "Segment", Template: "arn:${Partition}:evidently:${Region}:${Account}:segment/${SegmentName}"},
	})
}
