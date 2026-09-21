// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: panorama
// Source: https://servicereference.us-east-1.amazonaws.com/v1/panorama/panorama.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "panorama_application_instance", Service: "panorama", Resource: "applicationInstance", Template: "arn:${Partition}:panorama:${Region}:${Account}:applicationInstance/${ApplicationInstanceId}"},
		{Name: "panorama_device", Service: "panorama", Resource: "device", Template: "arn:${Partition}:panorama:${Region}:${Account}:device/${DeviceId}"},
		{Name: "panorama_package", Service: "panorama", Resource: "package", Template: "arn:${Partition}:panorama:${Region}:${Account}:package/${PackageId}"},
	})
}
