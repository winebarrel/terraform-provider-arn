// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: transfer
// Source: https://servicereference.us-east-1.amazonaws.com/v1/transfer/transfer.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "transfer_agreement", Service: "transfer", Resource: "agreement", Template: "arn:${Partition}:transfer:${Region}:${Account}:agreement/${ServerId}/${AgreementId}"},
		{Name: "transfer_certificate", Service: "transfer", Resource: "certificate", Template: "arn:${Partition}:transfer:${Region}:${Account}:certificate/${CertificateId}"},
		{Name: "transfer_connector", Service: "transfer", Resource: "connector", Template: "arn:${Partition}:transfer:${Region}:${Account}:connector/${ConnectorId}"},
		{Name: "transfer_host_key", Service: "transfer", Resource: "host-key", Template: "arn:${Partition}:transfer:${Region}:${Account}:host-key/${ServerId}/${HostKeyId}"},
		{Name: "transfer_profile", Service: "transfer", Resource: "profile", Template: "arn:${Partition}:transfer:${Region}:${Account}:profile/${ProfileId}"},
		{Name: "transfer_server", Service: "transfer", Resource: "server", Template: "arn:${Partition}:transfer:${Region}:${Account}:server/${ServerId}"},
		{Name: "transfer_user", Service: "transfer", Resource: "user", Template: "arn:${Partition}:transfer:${Region}:${Account}:user/${ServerId}/${UserName}"},
		{Name: "transfer_webapp", Service: "transfer", Resource: "webapp", Template: "arn:${Partition}:transfer:${Region}:${Account}:webapp/${WebAppId}"},
		{Name: "transfer_workflow", Service: "transfer", Resource: "workflow", Template: "arn:${Partition}:transfer:${Region}:${Account}:workflow/${WorkflowId}"},
	})
}
