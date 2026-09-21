// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: schemas
// Source: https://servicereference.us-east-1.amazonaws.com/v1/schemas/schemas.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "schemas_discoverer", Service: "schemas", Resource: "discoverer", Template: "arn:${Partition}:schemas:${Region}:${Account}:discoverer/${DiscovererId}"},
		{Name: "schemas_registry", Service: "schemas", Resource: "registry", Template: "arn:${Partition}:schemas:${Region}:${Account}:registry/${RegistryName}"},
		{Name: "schemas_schema", Service: "schemas", Resource: "schema", Template: "arn:${Partition}:schemas:${Region}:${Account}:schema/${RegistryName}/${SchemaName}"},
	})
}
