// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: s3vectors
// Source: https://servicereference.us-east-1.amazonaws.com/v1/s3vectors/s3vectors.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "s3vectors_index", Service: "s3vectors", Resource: "Index", Template: "arn:${Partition}:s3vectors:${Region}:${Account}:bucket/${BucketName}/index/${IndexName}"},
		{Name: "s3vectors_vector_bucket", Service: "s3vectors", Resource: "VectorBucket", Template: "arn:${Partition}:s3vectors:${Region}:${Account}:bucket/${BucketName}"},
	})
}
