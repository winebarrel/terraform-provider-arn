// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: databrew
// Source: https://servicereference.us-east-1.amazonaws.com/v1/databrew/databrew.json
// Functions: 6
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "databrew_dataset", Service: "databrew", Resource: "Dataset", Template: "arn:${Partition}:databrew:${Region}:${Account}:dataset/${ResourceId}"},
		{Name: "databrew_job", Service: "databrew", Resource: "Job", Template: "arn:${Partition}:databrew:${Region}:${Account}:job/${ResourceId}"},
		{Name: "databrew_project", Service: "databrew", Resource: "Project", Template: "arn:${Partition}:databrew:${Region}:${Account}:project/${ResourceId}"},
		{Name: "databrew_recipe", Service: "databrew", Resource: "Recipe", Template: "arn:${Partition}:databrew:${Region}:${Account}:recipe/${ResourceId}"},
		{Name: "databrew_ruleset", Service: "databrew", Resource: "Ruleset", Template: "arn:${Partition}:databrew:${Region}:${Account}:ruleset/${ResourceId}"},
		{Name: "databrew_schedule", Service: "databrew", Resource: "Schedule", Template: "arn:${Partition}:databrew:${Region}:${Account}:schedule/${ResourceId}"},
	})
}
