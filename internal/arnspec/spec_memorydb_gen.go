// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: memorydb
// Source: https://servicereference.us-east-1.amazonaws.com/v1/memorydb/memorydb.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "memorydb_acl", Service: "memorydb", Resource: "acl", Template: "arn:${Partition}:memorydb:${Region}:${Account}:acl/${AclName}"},
		{Name: "memorydb_cluster", Service: "memorydb", Resource: "cluster", Template: "arn:${Partition}:memorydb:${Region}:${Account}:cluster/${ClusterName}"},
		{Name: "memorydb_multiregioncluster", Service: "memorydb", Resource: "multiregioncluster", Template: "arn:${Partition}:memorydb::${Account}:multiregioncluster/${ClusterName}"},
		{Name: "memorydb_multiregionparametergroup", Service: "memorydb", Resource: "multiregionparametergroup", Template: "arn:${Partition}:memorydb::${Account}:multiregionparametergroup/${MultiRegionParameterGroupName}"},
		{Name: "memorydb_parametergroup", Service: "memorydb", Resource: "parametergroup", Template: "arn:${Partition}:memorydb:${Region}:${Account}:parametergroup/${ParameterGroupName}"},
		{Name: "memorydb_reservednode", Service: "memorydb", Resource: "reservednode", Template: "arn:${Partition}:memorydb:${Region}:${Account}:reservednode/${ReservationID}"},
		{Name: "memorydb_snapshot", Service: "memorydb", Resource: "snapshot", Template: "arn:${Partition}:memorydb:${Region}:${Account}:snapshot/${SnapshotName}"},
		{Name: "memorydb_subnetgroup", Service: "memorydb", Resource: "subnetgroup", Template: "arn:${Partition}:memorydb:${Region}:${Account}:subnetgroup/${SubnetGroupName}"},
		{Name: "memorydb_user", Service: "memorydb", Resource: "user", Template: "arn:${Partition}:memorydb:${Region}:${Account}:user/${UserName}"},
	})
}
