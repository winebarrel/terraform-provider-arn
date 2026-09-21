// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: odb
// Source: https://servicereference.us-east-1.amazonaws.com/v1/odb/odb.json
// Functions: 10
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "odb_autonomous_database", Service: "odb", Resource: "autonomous-database", Template: "arn:${Partition}:odb:${Region}:${Account}:autonomous-database/${AutonomousDatabaseId}"},
		{Name: "odb_autonomous_database_backup", Service: "odb", Resource: "autonomous-database-backup", Template: "arn:${Partition}:odb:${Region}:${Account}:autonomous-database-backup/${AutonomousDatabaseBackupId}"},
		{Name: "odb_cloud_autonomous_vm_cluster", Service: "odb", Resource: "cloud-autonomous-vm-cluster", Template: "arn:${Partition}:odb:${Region}:${Account}:cloud-autonomous-vm-cluster/${CloudAutonomousVmClusterId}"},
		{Name: "odb_cloud_exadata_infrastructure", Service: "odb", Resource: "cloud-exadata-infrastructure", Template: "arn:${Partition}:odb:${Region}:${Account}:cloud-exadata-infrastructure/${CloudExadataInfrastructureId}"},
		{Name: "odb_cloud_vm_cluster", Service: "odb", Resource: "cloud-vm-cluster", Template: "arn:${Partition}:odb:${Region}:${Account}:cloud-vm-cluster/${CloudVmClusterId}"},
		{Name: "odb_db_node", Service: "odb", Resource: "db-node", Template: "arn:${Partition}:odb:${Region}:${Account}:db-node/${DbNodeId}"},
		{Name: "odb_exadb_vm_cluster", Service: "odb", Resource: "exadb-vm-cluster", Template: "arn:${Partition}:odb:${Region}:${Account}:exadb-vm-cluster/${ExadbVmClusterId}"},
		{Name: "odb_exascale_db_storage_vault", Service: "odb", Resource: "exascale-db-storage-vault", Template: "arn:${Partition}:odb:${Region}:${Account}:exascale-db-storage-vault/${ExascaleDbStorageVaultId}"},
		{Name: "odb_odb_network", Service: "odb", Resource: "odb-network", Template: "arn:${Partition}:odb:${Region}:${Account}:odb-network/${OdbNetworkId}"},
		{Name: "odb_odb_peering_connection", Service: "odb", Resource: "odb-peering-connection", Template: "arn:${Partition}:odb:${Region}:${Account}:odb-peering-connection/${OdbPeeringConnectionId}"},
	})
}
