// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: s3
// Source: https://servicereference.us-east-1.amazonaws.com/v1/s3/s3.json
// Functions: 13
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "s3_accessgrant", Service: "s3", Resource: "accessgrant", Template: "arn:${Partition}:s3:${Region}:${Account}:access-grants/default/grant/${Token}"},
		{Name: "s3_accessgrantsinstance", Service: "s3", Resource: "accessgrantsinstance", Template: "arn:${Partition}:s3:${Region}:${Account}:access-grants/default"},
		{Name: "s3_accessgrantslocation", Service: "s3", Resource: "accessgrantslocation", Template: "arn:${Partition}:s3:${Region}:${Account}:access-grants/default/location/${Token}"},
		{Name: "s3_accesspoint", Service: "s3", Resource: "accesspoint", Template: "arn:${Partition}:s3:${Region}:${Account}:accesspoint/${AccessPointName}"},
		{Name: "s3_accesspointobject", Service: "s3", Resource: "accesspointobject", Template: "arn:${Partition}:s3:${Region}:${Account}:accesspoint/${AccessPointName}/object/${ObjectName}"},
		{Name: "s3_bucket", Service: "s3", Resource: "bucket", Template: "arn:${Partition}:s3:::${BucketName}"},
		{Name: "s3_job", Service: "s3", Resource: "job", Template: "arn:${Partition}:s3:${Region}:${Account}:job/${JobId}"},
		{Name: "s3_multiregionaccesspoint", Service: "s3", Resource: "multiregionaccesspoint", Template: "arn:${Partition}:s3::${Account}:accesspoint/${AccessPointAlias}"},
		{Name: "s3_multiregionaccesspointrequestarn", Service: "s3", Resource: "multiregionaccesspointrequestarn", Template: "arn:${Partition}:s3:us-west-2:${Account}:async-request/mrap/${Operation}/${Token}"},
		{Name: "s3_object", Service: "s3", Resource: "object", Template: "arn:${Partition}:s3:::${BucketName}/${ObjectName}"},
		{Name: "s3_objectlambdaaccesspoint", Service: "s3", Resource: "objectlambdaaccesspoint", Template: "arn:${Partition}:s3-object-lambda:${Region}:${Account}:accesspoint/${AccessPointName}"},
		{Name: "s3_storagelensconfiguration", Service: "s3", Resource: "storagelensconfiguration", Template: "arn:${Partition}:s3:${Region}:${Account}:storage-lens/${ConfigId}"},
		{Name: "s3_storagelensgroup", Service: "s3", Resource: "storagelensgroup", Template: "arn:${Partition}:s3:${Region}:${Account}:storage-lens-group/${Name}"},
	})
}
