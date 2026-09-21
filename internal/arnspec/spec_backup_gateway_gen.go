// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: backup-gateway
// Source: https://servicereference.us-east-1.amazonaws.com/v1/backup-gateway/backup-gateway.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "backup_gateway_gateway", Service: "backup-gateway", Resource: "gateway", Template: "arn:${Partition}:backup-gateway:${Region}:${Account}:gateway/${GatewayId}"},
		{Name: "backup_gateway_hypervisor", Service: "backup-gateway", Resource: "hypervisor", Template: "arn:${Partition}:backup-gateway:${Region}:${Account}:hypervisor/${HypervisorId}"},
		{Name: "backup_gateway_virtualmachine", Service: "backup-gateway", Resource: "virtualmachine", Template: "arn:${Partition}:backup-gateway:${Region}:${Account}:vm/${VirtualmachineId}"},
	})
}
