// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: batch
// Source: https://servicereference.us-east-1.amazonaws.com/v1/batch/batch.json
// Functions: 10
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "batch_compute_environment", Service: "batch", Resource: "compute-environment", Template: "arn:${Partition}:batch:${Region}:${Account}:compute-environment/${ComputeEnvironmentName}"},
		{Name: "batch_consumable_resource", Service: "batch", Resource: "consumable-resource", Template: "arn:${Partition}:batch:${Region}:${Account}:consumable-resource/${ConsumableResourceName}"},
		{Name: "batch_job", Service: "batch", Resource: "job", Template: "arn:${Partition}:batch:${Region}:${Account}:job/${JobId}"},
		{Name: "batch_job_definition", Service: "batch", Resource: "job-definition", Template: "arn:${Partition}:batch:${Region}:${Account}:job-definition/${JobDefinitionName}"},
		{Name: "batch_job_definition_revision", Service: "batch", Resource: "job-definition-revision", Template: "arn:${Partition}:batch:${Region}:${Account}:job-definition/${JobDefinitionName}:${Revision}"},
		{Name: "batch_job_queue", Service: "batch", Resource: "job-queue", Template: "arn:${Partition}:batch:${Region}:${Account}:job-queue/${JobQueueName}"},
		{Name: "batch_quota_share", Service: "batch", Resource: "quota-share", Template: "arn:${Partition}:batch:${Region}:${Account}:job-queue/${JobQueueName}/quota-share/${QuotaShareName}"},
		{Name: "batch_scheduling_policy", Service: "batch", Resource: "scheduling-policy", Template: "arn:${Partition}:batch:${Region}:${Account}:scheduling-policy/${SchedulingPolicyName}"},
		{Name: "batch_service_environment", Service: "batch", Resource: "service-environment", Template: "arn:${Partition}:batch:${Region}:${Account}:service-environment/${ServiceEnvironmentName}"},
		{Name: "batch_service_job", Service: "batch", Resource: "service-job", Template: "arn:${Partition}:batch:${Region}:${Account}:service-job/${JobId}"},
	})
}
