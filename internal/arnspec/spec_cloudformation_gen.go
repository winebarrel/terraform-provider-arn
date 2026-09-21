// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cloudformation
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cloudformation/cloudformation.json
// Functions: 8
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cloudformation_changeset", Service: "cloudformation", Resource: "changeset", Template: "arn:${Partition}:cloudformation:${Region}:${Account}:changeSet/${ChangeSetName}/${Id}"},
		{Name: "cloudformation_generatedtemplate", Service: "cloudformation", Resource: "generatedtemplate", Template: "arn:${Partition}:cloudformation:${Region}:${Account}:generatedTemplate/${Id}"},
		{Name: "cloudformation_resourcescan", Service: "cloudformation", Resource: "resourcescan", Template: "arn:${Partition}:cloudformation:${Region}:${Account}:resourceScan/${Id}"},
		{Name: "cloudformation_stack", Service: "cloudformation", Resource: "stack", Template: "arn:${Partition}:cloudformation:${Region}:${Account}:stack/${StackName}/${Id}"},
		{Name: "cloudformation_stackset", Service: "cloudformation", Resource: "stackset", Template: "arn:${Partition}:cloudformation:${Region}:${Account}:stackset/${StackSetName}:${Id}"},
		{Name: "cloudformation_stackset_target", Service: "cloudformation", Resource: "stackset-target", Template: "arn:${Partition}:cloudformation:${Region}:${Account}:stackset-target/${StackSetTarget}"},
		{Name: "cloudformation_type", Service: "cloudformation", Resource: "type", Template: "arn:${Partition}:cloudformation:${Region}:${Account}:type/resource/${Type}"},
		{Name: "cloudformation_type_hook", Service: "cloudformation", Resource: "typeHook", Template: "arn:${Partition}:cloudformation:${Region}:${Account}:type/hook/${Type}"},
	})
}
