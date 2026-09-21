// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: proton
// Source: https://servicereference.us-east-1.amazonaws.com/v1/proton/proton.json
// Functions: 15
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "proton_component", Service: "proton", Resource: "component", Template: "arn:${Partition}:proton:${Region}:${Account}:component/${Id}"},
		{Name: "proton_deployment", Service: "proton", Resource: "deployment", Template: "arn:${Partition}:proton:${Region}:${Account}:deployment/${Id}"},
		{Name: "proton_environment", Service: "proton", Resource: "environment", Template: "arn:${Partition}:proton:${Region}:${Account}:environment/${Name}"},
		{Name: "proton_environment_account_connection", Service: "proton", Resource: "environment-account-connection", Template: "arn:${Partition}:proton:${Region}:${Account}:environment-account-connection/${Id}"},
		{Name: "proton_environment_template", Service: "proton", Resource: "environment-template", Template: "arn:${Partition}:proton:${Region}:${Account}:environment-template/${Name}"},
		{Name: "proton_environment_template_major_version", Service: "proton", Resource: "environment-template-major-version", Template: "arn:${Partition}:proton:${Region}:${Account}:environment-template/${TemplateName}:${MajorVersionId}"},
		{Name: "proton_environment_template_minor_version", Service: "proton", Resource: "environment-template-minor-version", Template: "arn:${Partition}:proton:${Region}:${Account}:environment-template/${TemplateName}:${MajorVersionId}.${MinorVersionId}"},
		{Name: "proton_environment_template_version", Service: "proton", Resource: "environment-template-version", Template: "arn:${Partition}:proton:${Region}:${Account}:environment-template/${TemplateName}:${MajorVersion}.${MinorVersion}"},
		{Name: "proton_repository", Service: "proton", Resource: "repository", Template: "arn:${Partition}:proton:${Region}:${Account}:repository/${Provider}:${Name}"},
		{Name: "proton_service", Service: "proton", Resource: "service", Template: "arn:${Partition}:proton:${Region}:${Account}:service/${Name}"},
		{Name: "proton_service_instance", Service: "proton", Resource: "service-instance", Template: "arn:${Partition}:proton:${Region}:${Account}:service/${ServiceName}/service-instance/${Name}"},
		{Name: "proton_service_template", Service: "proton", Resource: "service-template", Template: "arn:${Partition}:proton:${Region}:${Account}:service-template/${Name}"},
		{Name: "proton_service_template_major_version", Service: "proton", Resource: "service-template-major-version", Template: "arn:${Partition}:proton:${Region}:${Account}:service-template/${TemplateName}:${MajorVersionId}"},
		{Name: "proton_service_template_minor_version", Service: "proton", Resource: "service-template-minor-version", Template: "arn:${Partition}:proton:${Region}:${Account}:service-template/${TemplateName}:${MajorVersionId}.${MinorVersionId}"},
		{Name: "proton_service_template_version", Service: "proton", Resource: "service-template-version", Template: "arn:${Partition}:proton:${Region}:${Account}:service-template/${TemplateName}:${MajorVersion}.${MinorVersion}"},
	})
}
