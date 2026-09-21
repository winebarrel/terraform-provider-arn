// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: dms
// Source: https://servicereference.us-east-1.amazonaws.com/v1/dms/dms.json
// Functions: 13
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "dms_certificate", Service: "dms", Resource: "Certificate", Template: "arn:${Partition}:dms:${Region}:${Account}:cert:*"},
		{Name: "dms_data_migration", Service: "dms", Resource: "DataMigration", Template: "arn:${Partition}:dms:${Region}:${Account}:data-migration:*"},
		{Name: "dms_data_provider", Service: "dms", Resource: "DataProvider", Template: "arn:${Partition}:dms:${Region}:${Account}:data-provider:*"},
		{Name: "dms_endpoint", Service: "dms", Resource: "Endpoint", Template: "arn:${Partition}:dms:${Region}:${Account}:endpoint:*"},
		{Name: "dms_event_subscription", Service: "dms", Resource: "EventSubscription", Template: "arn:${Partition}:dms:${Region}:${Account}:es:*"},
		{Name: "dms_instance_profile", Service: "dms", Resource: "InstanceProfile", Template: "arn:${Partition}:dms:${Region}:${Account}:instance-profile:*"},
		{Name: "dms_migration_project", Service: "dms", Resource: "MigrationProject", Template: "arn:${Partition}:dms:${Region}:${Account}:migration-project:*"},
		{Name: "dms_replication_config", Service: "dms", Resource: "ReplicationConfig", Template: "arn:${Partition}:dms:${Region}:${Account}:replication-config:*"},
		{Name: "dms_replication_instance", Service: "dms", Resource: "ReplicationInstance", Template: "arn:${Partition}:dms:${Region}:${Account}:rep:*"},
		{Name: "dms_replication_subnet_group", Service: "dms", Resource: "ReplicationSubnetGroup", Template: "arn:${Partition}:dms:${Region}:${Account}:subgrp:*"},
		{Name: "dms_replication_task", Service: "dms", Resource: "ReplicationTask", Template: "arn:${Partition}:dms:${Region}:${Account}:task:*"},
		{Name: "dms_replication_task_assessment_run", Service: "dms", Resource: "ReplicationTaskAssessmentRun", Template: "arn:${Partition}:dms:${Region}:${Account}:assessment-run:*"},
		{Name: "dms_replication_task_individual_assessment", Service: "dms", Resource: "ReplicationTaskIndividualAssessment", Template: "arn:${Partition}:dms:${Region}:${Account}:individual-assessment:*"},
	})
}
