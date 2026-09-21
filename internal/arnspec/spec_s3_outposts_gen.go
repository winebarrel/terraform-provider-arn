// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: s3-outposts
// Source: https://servicereference.us-east-1.amazonaws.com/v1/s3-outposts/s3-outposts.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "s3_outposts_accesspoint", Service: "s3-outposts", Resource: "accesspoint", Template: "arn:${Partition}:s3-outposts:${Region}:${Account}:outpost/${OutpostId}/accesspoint/${AccessPointName}"},
		{Name: "s3_outposts_bucket", Service: "s3-outposts", Resource: "bucket", Template: "arn:${Partition}:s3-outposts:${Region}:${Account}:outpost/${OutpostId}/bucket/${BucketName}"},
		{Name: "s3_outposts_endpoint", Service: "s3-outposts", Resource: "endpoint", Template: "arn:${Partition}:s3-outposts:${Region}:${Account}:outpost/${OutpostId}/endpoint/${EndpointId}"},
		{Name: "s3_outposts_object", Service: "s3-outposts", Resource: "object", Template: "arn:${Partition}:s3-outposts:${Region}:${Account}:outpost/${OutpostId}/bucket/${BucketName}/object/${ObjectName}"},
	})
}
