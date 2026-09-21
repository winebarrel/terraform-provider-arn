// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: clouddirectory
// Source: https://servicereference.us-east-1.amazonaws.com/v1/clouddirectory/clouddirectory.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "clouddirectory_applied_schema", Service: "clouddirectory", Resource: "appliedSchema", Template: "arn:${Partition}:clouddirectory:${Region}:${Account}:directory/${DirectoryId}/schema/${SchemaName}/${Version}"},
		{Name: "clouddirectory_development_schema", Service: "clouddirectory", Resource: "developmentSchema", Template: "arn:${Partition}:clouddirectory:${Region}:${Account}:schema/development/${SchemaName}"},
		{Name: "clouddirectory_directory", Service: "clouddirectory", Resource: "directory", Template: "arn:${Partition}:clouddirectory:${Region}:${Account}:directory/${DirectoryId}"},
		{Name: "clouddirectory_published_schema", Service: "clouddirectory", Resource: "publishedSchema", Template: "arn:${Partition}:clouddirectory:${Region}:${Account}:schema/published/${SchemaName}/${Version}"},
	})
}
