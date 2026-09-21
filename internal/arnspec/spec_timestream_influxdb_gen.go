// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: timestream-influxdb
// Source: https://servicereference.us-east-1.amazonaws.com/v1/timestream-influxdb/timestream-influxdb.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "timestream_influxdb_db_backup", Service: "timestream-influxdb", Resource: "db-backup", Template: "arn:${Partition}:timestream-influxdb:${Region}:${Account}:db-backup/${DbBackupId}"},
		{Name: "timestream_influxdb_db_cluster", Service: "timestream-influxdb", Resource: "db-cluster", Template: "arn:${Partition}:timestream-influxdb:${Region}:${Account}:db-cluster/${DbClusterId}"},
		{Name: "timestream_influxdb_db_instance", Service: "timestream-influxdb", Resource: "db-instance", Template: "arn:${Partition}:timestream-influxdb:${Region}:${Account}:db-instance/${DbInstanceIdentifier}"},
		{Name: "timestream_influxdb_db_parameter_group", Service: "timestream-influxdb", Resource: "db-parameter-group", Template: "arn:${Partition}:timestream-influxdb:${Region}:${Account}:db-parameter-group/${DbParameterGroupIdentifier}"},
	})
}
