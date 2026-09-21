// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: rds
// Source: https://servicereference.us-east-1.amazonaws.com/v1/rds/rds.json
// Functions: 24
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "rds_auto_backup", Service: "rds", Resource: "auto-backup", Template: "arn:${Partition}:rds:${Region}:${Account}:auto-backup:${DbInstanceAutomatedBackupId}"},
		{Name: "rds_cev", Service: "rds", Resource: "cev", Template: "arn:${Partition}:rds:${Region}:${Account}:cev:${Engine}/${EngineVersion}/${CustomDbEngineVersionId}"},
		{Name: "rds_cluster", Service: "rds", Resource: "cluster", Template: "arn:${Partition}:rds:${Region}:${Account}:cluster:${DbClusterInstanceName}"},
		{Name: "rds_cluster_auto_backup", Service: "rds", Resource: "cluster-auto-backup", Template: "arn:${Partition}:rds:${Region}:${Account}:cluster-auto-backup:${DbClusterAutomatedBackupId}"},
		{Name: "rds_cluster_endpoint", Service: "rds", Resource: "cluster-endpoint", Template: "arn:${Partition}:rds:${Region}:${Account}:cluster-endpoint:${DbClusterEndpoint}"},
		{Name: "rds_cluster_pg", Service: "rds", Resource: "cluster-pg", Template: "arn:${Partition}:rds:${Region}:${Account}:cluster-pg:${ClusterParameterGroupName}"},
		{Name: "rds_cluster_snapshot", Service: "rds", Resource: "cluster-snapshot", Template: "arn:${Partition}:rds:${Region}:${Account}:cluster-snapshot:${ClusterSnapshotName}"},
		{Name: "rds_db", Service: "rds", Resource: "db", Template: "arn:${Partition}:rds:${Region}:${Account}:db:${DbInstanceName}"},
		{Name: "rds_deployment", Service: "rds", Resource: "deployment", Template: "arn:${Partition}:rds:${Region}:${Account}:deployment:${BlueGreenDeploymentIdentifier}"},
		{Name: "rds_es", Service: "rds", Resource: "es", Template: "arn:${Partition}:rds:${Region}:${Account}:es:${SubscriptionName}"},
		{Name: "rds_global_cluster", Service: "rds", Resource: "global-cluster", Template: "arn:${Partition}:rds::${Account}:global-cluster:${GlobalCluster}"},
		{Name: "rds_integration", Service: "rds", Resource: "integration", Template: "arn:${Partition}:rds:${Region}:${Account}:integration:${IntegrationIdentifier}"},
		{Name: "rds_og", Service: "rds", Resource: "og", Template: "arn:${Partition}:rds:${Region}:${Account}:og:${OptionGroupName}"},
		{Name: "rds_pg", Service: "rds", Resource: "pg", Template: "arn:${Partition}:rds:${Region}:${Account}:pg:${ParameterGroupName}"},
		{Name: "rds_proxy", Service: "rds", Resource: "proxy", Template: "arn:${Partition}:rds:${Region}:${Account}:db-proxy:${DbProxyId}"},
		{Name: "rds_proxy_endpoint", Service: "rds", Resource: "proxy-endpoint", Template: "arn:${Partition}:rds:${Region}:${Account}:db-proxy-endpoint:${DbProxyEndpointId}"},
		{Name: "rds_ri", Service: "rds", Resource: "ri", Template: "arn:${Partition}:rds:${Region}:${Account}:ri:${ReservedDbInstanceName}"},
		{Name: "rds_secgrp", Service: "rds", Resource: "secgrp", Template: "arn:${Partition}:rds:${Region}:${Account}:secgrp:${SecurityGroupName}"},
		{Name: "rds_shardgrp", Service: "rds", Resource: "shardgrp", Template: "arn:${Partition}:rds:${Region}:${Account}:shard-group:${DbShardGroupResourceId}"},
		{Name: "rds_snapshot", Service: "rds", Resource: "snapshot", Template: "arn:${Partition}:rds:${Region}:${Account}:snapshot:${SnapshotName}"},
		{Name: "rds_snapshot_tenant_database", Service: "rds", Resource: "snapshot-tenant-database", Template: "arn:${Partition}:rds:${Region}:${Account}:snapshot-tenant-database:${SnapshotName}:${TenantResourceId}"},
		{Name: "rds_subgrp", Service: "rds", Resource: "subgrp", Template: "arn:${Partition}:rds:${Region}:${Account}:subgrp:${SubnetGroupName}"},
		{Name: "rds_target_group", Service: "rds", Resource: "target-group", Template: "arn:${Partition}:rds:${Region}:${Account}:target-group:${TargetGroupId}"},
		{Name: "rds_tenant_database", Service: "rds", Resource: "tenant-database", Template: "arn:${Partition}:rds:${Region}:${Account}:tenant-database:${TenantResourceId}"},
	})
}
