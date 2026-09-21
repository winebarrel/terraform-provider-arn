// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: mediastore
// Source: https://servicereference.us-east-1.amazonaws.com/v1/mediastore/mediastore.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "mediastore_container", Service: "mediastore", Resource: "container", Template: "arn:${Partition}:mediastore:${Region}:${Account}:container/${ContainerName}"},
		{Name: "mediastore_folder", Service: "mediastore", Resource: "folder", Template: "arn:${Partition}:mediastore:${Region}:${Account}:container/${ContainerName}/${FolderPath}"},
		{Name: "mediastore_object", Service: "mediastore", Resource: "object", Template: "arn:${Partition}:mediastore:${Region}:${Account}:container/${ContainerName}/${ObjectPath}"},
	})
}
