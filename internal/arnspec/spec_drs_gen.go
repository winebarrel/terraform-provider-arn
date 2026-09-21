// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: drs
// Source: https://servicereference.us-east-1.amazonaws.com/v1/drs/drs.json
// Functions: 8
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "drs_job_resource", Service: "drs", Resource: "JobResource", Template: "arn:${Partition}:drs:${Region}:${Account}:job/${JobID}"},
		{Name: "drs_launch_configuration_template_resource", Service: "drs", Resource: "LaunchConfigurationTemplateResource", Template: "arn:${Partition}:drs:${Region}:${Account}:launch-configuration-template/${LaunchConfigurationTemplateID}"},
		{Name: "drs_recovery_instance_resource", Service: "drs", Resource: "RecoveryInstanceResource", Template: "arn:${Partition}:drs:${Region}:${Account}:recovery-instance/${RecoveryInstanceID}"},
		{Name: "drs_recovery_plan_execution_resource", Service: "drs", Resource: "RecoveryPlanExecutionResource", Template: "arn:${Partition}:drs:${Region}:${Account}:recovery-plan-execution/${RecoveryPlanExecutionID}"},
		{Name: "drs_recovery_plan_resource", Service: "drs", Resource: "RecoveryPlanResource", Template: "arn:${Partition}:drs:${Region}:${Account}:recovery-plan/${RecoveryPlanID}"},
		{Name: "drs_replication_configuration_template_resource", Service: "drs", Resource: "ReplicationConfigurationTemplateResource", Template: "arn:${Partition}:drs:${Region}:${Account}:replication-configuration-template/${ReplicationConfigurationTemplateID}"},
		{Name: "drs_source_network_resource", Service: "drs", Resource: "SourceNetworkResource", Template: "arn:${Partition}:drs:${Region}:${Account}:source-network/${SourceNetworkID}"},
		{Name: "drs_source_server_resource", Service: "drs", Resource: "SourceServerResource", Template: "arn:${Partition}:drs:${Region}:${Account}:source-server/${SourceServerID}"},
	})
}
