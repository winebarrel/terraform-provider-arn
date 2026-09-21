// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ds
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ds/ds.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ds_directory", Service: "ds", Resource: "directory", Template: "arn:${Partition}:ds:${Region}:${Account}:directory/${DirectoryId}"},
	})
}
