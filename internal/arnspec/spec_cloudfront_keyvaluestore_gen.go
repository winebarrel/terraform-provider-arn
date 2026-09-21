// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cloudfront-keyvaluestore
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cloudfront-keyvaluestore/cloudfront-keyvaluestore.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cloudfront_keyvaluestore_key_value_store", Service: "cloudfront-keyvaluestore", Resource: "key-value-store", Template: "arn:${Partition}:cloudfront::${Account}:key-value-store/${ResourceId}"},
	})
}
