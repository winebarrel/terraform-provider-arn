// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: devicefarm
// Source: https://servicereference.us-east-1.amazonaws.com/v1/devicefarm/devicefarm.json
// Functions: 17
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "devicefarm_artifact", Service: "devicefarm", Resource: "artifact", Template: "arn:${Partition}:devicefarm:${Region}:${Account}:artifact:${ResourceId}"},
		{Name: "devicefarm_device", Service: "devicefarm", Resource: "device", Template: "arn:${Partition}:devicefarm:${Region}::device:${ResourceId}"},
		{Name: "devicefarm_deviceinstance", Service: "devicefarm", Resource: "deviceinstance", Template: "arn:${Partition}:devicefarm:${Region}::deviceinstance:${ResourceId}"},
		{Name: "devicefarm_devicepool", Service: "devicefarm", Resource: "devicepool", Template: "arn:${Partition}:devicefarm:${Region}:${Account}:devicepool:${ResourceId}"},
		{Name: "devicefarm_instanceprofile", Service: "devicefarm", Resource: "instanceprofile", Template: "arn:${Partition}:devicefarm:${Region}:${Account}:instanceprofile:${ResourceId}"},
		{Name: "devicefarm_job", Service: "devicefarm", Resource: "job", Template: "arn:${Partition}:devicefarm:${Region}:${Account}:job:${ResourceId}"},
		{Name: "devicefarm_networkprofile", Service: "devicefarm", Resource: "networkprofile", Template: "arn:${Partition}:devicefarm:${Region}:${Account}:networkprofile:${ResourceId}"},
		{Name: "devicefarm_project", Service: "devicefarm", Resource: "project", Template: "arn:${Partition}:devicefarm:${Region}:${Account}:project:${ResourceId}"},
		{Name: "devicefarm_run", Service: "devicefarm", Resource: "run", Template: "arn:${Partition}:devicefarm:${Region}:${Account}:run:${ResourceId}"},
		{Name: "devicefarm_sample", Service: "devicefarm", Resource: "sample", Template: "arn:${Partition}:devicefarm:${Region}:${Account}:sample:${ResourceId}"},
		{Name: "devicefarm_session", Service: "devicefarm", Resource: "session", Template: "arn:${Partition}:devicefarm:${Region}:${Account}:session:${ResourceId}"},
		{Name: "devicefarm_suite", Service: "devicefarm", Resource: "suite", Template: "arn:${Partition}:devicefarm:${Region}:${Account}:suite:${ResourceId}"},
		{Name: "devicefarm_test", Service: "devicefarm", Resource: "test", Template: "arn:${Partition}:devicefarm:${Region}:${Account}:test:${ResourceId}"},
		{Name: "devicefarm_testgrid_project", Service: "devicefarm", Resource: "testgrid-project", Template: "arn:${Partition}:devicefarm:${Region}:${Account}:testgrid-project:${ResourceId}"},
		{Name: "devicefarm_testgrid_session", Service: "devicefarm", Resource: "testgrid-session", Template: "arn:${Partition}:devicefarm:${Region}:${Account}:testgrid-session:${ResourceId}"},
		{Name: "devicefarm_upload", Service: "devicefarm", Resource: "upload", Template: "arn:${Partition}:devicefarm:${Region}:${Account}:upload:${ResourceId}"},
		{Name: "devicefarm_vpceconfiguration", Service: "devicefarm", Resource: "vpceconfiguration", Template: "arn:${Partition}:devicefarm:${Region}:${Account}:vpceconfiguration:${ResourceId}"},
	})
}
