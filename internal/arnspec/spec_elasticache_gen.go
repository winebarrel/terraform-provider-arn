// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: elasticache
// Source: https://servicereference.us-east-1.amazonaws.com/v1/elasticache/elasticache.json
// Functions: 12
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "elasticache_cluster", Service: "elasticache", Resource: "cluster", Template: "arn:${Partition}:elasticache:${Region}:${Account}:cluster:${CacheClusterId}"},
		{Name: "elasticache_globalreplicationgroup", Service: "elasticache", Resource: "globalreplicationgroup", Template: "arn:${Partition}:elasticache::${Account}:globalreplicationgroup:${GlobalReplicationGroupId}"},
		{Name: "elasticache_parametergroup", Service: "elasticache", Resource: "parametergroup", Template: "arn:${Partition}:elasticache:${Region}:${Account}:parametergroup:${CacheParameterGroupName}"},
		{Name: "elasticache_replicationgroup", Service: "elasticache", Resource: "replicationgroup", Template: "arn:${Partition}:elasticache:${Region}:${Account}:replicationgroup:${ReplicationGroupId}"},
		{Name: "elasticache_reserved_instance", Service: "elasticache", Resource: "reserved-instance", Template: "arn:${Partition}:elasticache:${Region}:${Account}:reserved-instance:${ReservedCacheNodeId}"},
		{Name: "elasticache_securitygroup", Service: "elasticache", Resource: "securitygroup", Template: "arn:${Partition}:elasticache:${Region}:${Account}:securitygroup:${CacheSecurityGroupName}"},
		{Name: "elasticache_serverlesscache", Service: "elasticache", Resource: "serverlesscache", Template: "arn:${Partition}:elasticache:${Region}:${Account}:serverlesscache:${ServerlessCacheName}"},
		{Name: "elasticache_serverlesscachesnapshot", Service: "elasticache", Resource: "serverlesscachesnapshot", Template: "arn:${Partition}:elasticache:${Region}:${Account}:serverlesscachesnapshot:${ServerlessCacheSnapshotName}"},
		{Name: "elasticache_snapshot", Service: "elasticache", Resource: "snapshot", Template: "arn:${Partition}:elasticache:${Region}:${Account}:snapshot:${SnapshotName}"},
		{Name: "elasticache_subnetgroup", Service: "elasticache", Resource: "subnetgroup", Template: "arn:${Partition}:elasticache:${Region}:${Account}:subnetgroup:${CacheSubnetGroupName}"},
		{Name: "elasticache_user", Service: "elasticache", Resource: "user", Template: "arn:${Partition}:elasticache:${Region}:${Account}:user:${UserId}"},
		{Name: "elasticache_usergroup", Service: "elasticache", Resource: "usergroup", Template: "arn:${Partition}:elasticache:${Region}:${Account}:usergroup:${UserGroupId}"},
	})
}
