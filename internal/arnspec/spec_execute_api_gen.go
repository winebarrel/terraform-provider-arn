// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: execute-api
// Source: https://servicereference.us-east-1.amazonaws.com/v1/execute-api/execute-api.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "execute_api_execute_api_domain", Service: "execute-api", Resource: "execute-api-domain", Template: "arn:${Partition}:execute-api:${Region}:${Account}:/domainnames/${DomainName}+${DomainIdentifier}"},
		{Name: "execute_api_execute_api_general", Service: "execute-api", Resource: "execute-api-general", Template: "arn:${Partition}:execute-api:${Region}:${Account}:${ApiId}/${Stage}/${Method}/${ApiSpecificResourcePath}"},
	})
}
