// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: neptune-db
// Source: https://servicereference.us-east-1.amazonaws.com/v1/neptune-db/neptune-db.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "neptune_db_database", Service: "neptune-db", Resource: "database", Template: "arn:${Partition}:neptune-db:${Region}:${Account}:${ClusterResourceId}/*"},
	})
}
