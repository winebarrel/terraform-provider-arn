// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: amplifyuibuilder
// Source: https://servicereference.us-east-1.amazonaws.com/v1/amplifyuibuilder/amplifyuibuilder.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "amplifyuibuilder_codegen_job_resource", Service: "amplifyuibuilder", Resource: "CodegenJobResource", Template: "arn:${Partition}:amplifyuibuilder:${Region}:${Account}:app/${AppId}/environment/${EnvironmentName}/codegen-jobs/${Id}"},
		{Name: "amplifyuibuilder_component_resource", Service: "amplifyuibuilder", Resource: "ComponentResource", Template: "arn:${Partition}:amplifyuibuilder:${Region}:${Account}:app/${AppId}/environment/${EnvironmentName}/components/${Id}"},
		{Name: "amplifyuibuilder_form_resource", Service: "amplifyuibuilder", Resource: "FormResource", Template: "arn:${Partition}:amplifyuibuilder:${Region}:${Account}:app/${AppId}/environment/${EnvironmentName}/forms/${Id}"},
		{Name: "amplifyuibuilder_theme_resource", Service: "amplifyuibuilder", Resource: "ThemeResource", Template: "arn:${Partition}:amplifyuibuilder:${Region}:${Account}:app/${AppId}/environment/${EnvironmentName}/themes/${Id}"},
	})
}
