// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: amplify
// Source: https://servicereference.us-east-1.amazonaws.com/v1/amplify/amplify.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "amplify_apps", Service: "amplify", Resource: "apps", Template: "arn:${Partition}:amplify:${Region}:${Account}:apps/${AppId}"},
		{Name: "amplify_branches", Service: "amplify", Resource: "branches", Template: "arn:${Partition}:amplify:${Region}:${Account}:apps/${AppId}/branches/${BranchName}"},
		{Name: "amplify_domains", Service: "amplify", Resource: "domains", Template: "arn:${Partition}:amplify:${Region}:${Account}:apps/${AppId}/domains/${DomainName}"},
		{Name: "amplify_jobs", Service: "amplify", Resource: "jobs", Template: "arn:${Partition}:amplify:${Region}:${Account}:apps/${AppId}/branches/${BranchName}/jobs/${JobId}"},
		{Name: "amplify_webhooks", Service: "amplify", Resource: "webhooks", Template: "arn:${Partition}:amplify:${Region}:${Account}:webhooks/${WebhookId}"},
	})
}
