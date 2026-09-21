// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: elasticfilesystem
// Source: https://servicereference.us-east-1.amazonaws.com/v1/elasticfilesystem/elasticfilesystem.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "elasticfilesystem_access_point", Service: "elasticfilesystem", Resource: "access-point", Template: "arn:${Partition}:elasticfilesystem:${Region}:${Account}:access-point/${AccessPointId}"},
		{Name: "elasticfilesystem_file_system", Service: "elasticfilesystem", Resource: "file-system", Template: "arn:${Partition}:elasticfilesystem:${Region}:${Account}:file-system/${FileSystemId}"},
	})
}
