// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: refactor-spaces
// Source: https://servicereference.us-east-1.amazonaws.com/v1/refactor-spaces/refactor-spaces.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "refactor_spaces_application", Service: "refactor-spaces", Resource: "application", Template: "arn:${Partition}:refactor-spaces:${Region}:${Account}:environment/${EnvironmentId}/application/${ApplicationId}"},
		{Name: "refactor_spaces_environment", Service: "refactor-spaces", Resource: "environment", Template: "arn:${Partition}:refactor-spaces:${Region}:${Account}:environment/${EnvironmentId}"},
		{Name: "refactor_spaces_route", Service: "refactor-spaces", Resource: "route", Template: "arn:${Partition}:refactor-spaces:${Region}:${Account}:environment/${EnvironmentId}/application/${ApplicationId}/route/${RouteId}"},
		{Name: "refactor_spaces_service", Service: "refactor-spaces", Resource: "service", Template: "arn:${Partition}:refactor-spaces:${Region}:${Account}:environment/${EnvironmentId}/application/${ApplicationId}/service/${ServiceId}"},
	})
}
