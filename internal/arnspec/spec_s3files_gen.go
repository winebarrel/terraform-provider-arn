// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: s3files
// Source: https://servicereference.us-east-1.amazonaws.com/v1/s3files/s3files.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "s3files_access_point", Service: "s3files", Resource: "access-point", Template: "arn:${Partition}:s3files:${Region}:${Account}:file-system/${FileSystemId}/access-point/${AccessPointId}"},
		{Name: "s3files_file_system", Service: "s3files", Resource: "file-system", Template: "arn:${Partition}:s3files:${Region}:${Account}:file-system/${FileSystemId}"},
	})
}
