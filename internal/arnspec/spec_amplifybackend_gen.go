// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: amplifybackend
// Source: https://servicereference.us-east-1.amazonaws.com/v1/amplifybackend/amplifybackend.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "amplifybackend_api", Service: "amplifybackend", Resource: "api", Template: "arn:${Partition}:amplifybackend:${Region}:${Account}:/backend/${AppId}/api/*"},
		{Name: "amplifybackend_auth", Service: "amplifybackend", Resource: "auth", Template: "arn:${Partition}:amplifybackend:${Region}:${Account}:/backend/${AppId}/auth/*"},
		{Name: "amplifybackend_backend", Service: "amplifybackend", Resource: "backend", Template: "arn:${Partition}:amplifybackend:${Region}:${Account}:/backend/${AppId}/*"},
		{Name: "amplifybackend_config", Service: "amplifybackend", Resource: "config", Template: "arn:${Partition}:amplifybackend:${Region}:${Account}:/backend/${AppId}/config/*"},
		{Name: "amplifybackend_created_backend", Service: "amplifybackend", Resource: "created-backend", Template: "arn:${Partition}:amplifybackend:${Region}:${Account}:/backend/*"},
		{Name: "amplifybackend_environment", Service: "amplifybackend", Resource: "environment", Template: "arn:${Partition}:amplifybackend:${Region}:${Account}:/backend/${AppId}/environments/*"},
		{Name: "amplifybackend_job", Service: "amplifybackend", Resource: "job", Template: "arn:${Partition}:amplifybackend:${Region}:${Account}:/backend/${AppId}/job/*"},
		{Name: "amplifybackend_storage", Service: "amplifybackend", Resource: "storage", Template: "arn:${Partition}:amplifybackend:${Region}:${Account}:/backend/${AppId}/storage/*"},
		{Name: "amplifybackend_token", Service: "amplifybackend", Resource: "token", Template: "arn:${Partition}:amplifybackend:${Region}:${Account}:/backend/${AppId}/challenge/*"},
	})
}
