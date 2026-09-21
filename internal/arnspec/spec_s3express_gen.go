// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: s3express
// Source: https://servicereference.us-east-1.amazonaws.com/v1/s3express/s3express.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "s3express_accesspoint", Service: "s3express", Resource: "accesspoint", Template: "arn:${Partition}:s3express:${Region}:${Account}:accesspoint/${AccessPointName}"},
		{Name: "s3express_bucket", Service: "s3express", Resource: "bucket", Template: "arn:${Partition}:s3express:${Region}:${Account}:bucket/${BucketName}"},
	})
}
