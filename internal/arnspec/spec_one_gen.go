// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: one
// Source: https://servicereference.us-east-1.amazonaws.com/v1/one/one.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "one_configuration", Service: "one", Resource: "configuration", Template: "arn:${Partition}:one:${Region}:${Account}:device-instance/${DeviceInstanceId}/configuration/${Version}"},
		{Name: "one_device_configuration_template", Service: "one", Resource: "device-configuration-template", Template: "arn:${Partition}:one:${Region}:${Account}:device-configuration-template/${TemplateId}"},
		{Name: "one_device_instance", Service: "one", Resource: "device-instance", Template: "arn:${Partition}:one:${Region}:${Account}:device-instance/${DeviceInstanceId}"},
		{Name: "one_site", Service: "one", Resource: "site", Template: "arn:${Partition}:one:${Region}:${Account}:site/${SiteId}"},
		{Name: "one_user", Service: "one", Resource: "user", Template: "arn:${Partition}:one:${Region}:${Account}:user/${UserId}"},
	})
}
