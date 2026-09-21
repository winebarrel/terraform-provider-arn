// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: rds-db
// Source: https://servicereference.us-east-1.amazonaws.com/v1/rds-db/rds-db.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "rds_db_db_user", Service: "rds-db", Resource: "db-user", Template: "arn:${Partition}:rds-db:${Region}:${Account}:dbuser:${DbiResourceId}/${DbUserName}"},
	})
}
