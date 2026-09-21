// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: s3-object-lambda
// Source: https://servicereference.us-east-1.amazonaws.com/v1/s3-object-lambda/s3-object-lambda.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "s3_object_lambda_objectlambdaaccesspoint", Service: "s3-object-lambda", Resource: "objectlambdaaccesspoint", Template: "arn:${Partition}:s3-object-lambda:${Region}:${Account}:accesspoint/${AccessPointName}"},
	})
}
