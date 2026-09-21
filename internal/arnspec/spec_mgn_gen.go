// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: mgn
// Source: https://servicereference.us-east-1.amazonaws.com/v1/mgn/mgn.json
// Functions: 11
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "mgn_application_resource", Service: "mgn", Resource: "ApplicationResource", Template: "arn:${Partition}:mgn:${Region}:${Account}:application/${ApplicationID}"},
		{Name: "mgn_connector_resource", Service: "mgn", Resource: "ConnectorResource", Template: "arn:${Partition}:mgn:${Region}:${Account}:connector/${ConnectorID}"},
		{Name: "mgn_export_resource", Service: "mgn", Resource: "ExportResource", Template: "arn:${Partition}:mgn:${Region}:${Account}:export/${ExportID}"},
		{Name: "mgn_import_resource", Service: "mgn", Resource: "ImportResource", Template: "arn:${Partition}:mgn:${Region}:${Account}:import/${ImportID}"},
		{Name: "mgn_job_resource", Service: "mgn", Resource: "JobResource", Template: "arn:${Partition}:mgn:${Region}:${Account}:job/${JobID}"},
		{Name: "mgn_launch_configuration_template_resource", Service: "mgn", Resource: "LaunchConfigurationTemplateResource", Template: "arn:${Partition}:mgn:${Region}:${Account}:launch-configuration-template/${LaunchConfigurationTemplateID}"},
		{Name: "mgn_network_migration_definition_resource", Service: "mgn", Resource: "NetworkMigrationDefinitionResource", Template: "arn:${Partition}:mgn:${Region}:${Account}:network-migration-definition/${NetworkMigrationDefinitionID}"},
		{Name: "mgn_replication_configuration_template_resource", Service: "mgn", Resource: "ReplicationConfigurationTemplateResource", Template: "arn:${Partition}:mgn:${Region}:${Account}:replication-configuration-template/${ReplicationConfigurationTemplateID}"},
		{Name: "mgn_source_server_resource", Service: "mgn", Resource: "SourceServerResource", Template: "arn:${Partition}:mgn:${Region}:${Account}:source-server/${SourceServerID}"},
		{Name: "mgn_vcenter_client_resource", Service: "mgn", Resource: "VcenterClientResource", Template: "arn:${Partition}:mgn:${Region}:${Account}:vcenter-client/${VcenterClientID}"},
		{Name: "mgn_wave_resource", Service: "mgn", Resource: "WaveResource", Template: "arn:${Partition}:mgn:${Region}:${Account}:wave/${WaveID}"},
	})
}
