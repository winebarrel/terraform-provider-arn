// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: signin
// Source: https://servicereference.us-east-1.amazonaws.com/v1/signin/signin.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "signin_console", Service: "signin", Resource: "console", Template: "arn:${Partition}:signin:::console/${ConsoleName}"},
		{Name: "signin_oauth2_public_client_localhost", Service: "signin", Resource: "oauth2-public-client-localhost", Template: "arn:${Partition}:signin:${Region}:${Account}:oauth2/public-client/localhost"},
		{Name: "signin_oauth2_public_client_registration", Service: "signin", Resource: "oauth2-public-client-registration", Template: "arn:${Partition}:signin:${Region}::external-client/dcr/*"},
		{Name: "signin_oauth2_public_client_remote", Service: "signin", Resource: "oauth2-public-client-remote", Template: "arn:${Partition}:signin:${Region}:${Account}:oauth2/public-client/remote"},
		{Name: "signin_oauth2_resource_service_principal", Service: "signin", Resource: "oauth2-resource-service-principal", Template: "arn:${Partition}:signin:${Region}:${Account}:service-principal/${ServicePrincipalName}"},
	})
}
