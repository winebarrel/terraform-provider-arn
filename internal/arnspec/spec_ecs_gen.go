// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ecs
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ecs/ecs.json
// Functions: 13
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ecs_capacity_provider", Service: "ecs", Resource: "capacity-provider", Template: "arn:${Partition}:ecs:${Region}:${Account}:capacity-provider/${CapacityProviderName}"},
		{Name: "ecs_cluster", Service: "ecs", Resource: "cluster", Template: "arn:${Partition}:ecs:${Region}:${Account}:cluster/${ClusterName}"},
		{Name: "ecs_container_instance", Service: "ecs", Resource: "container-instance", Template: "arn:${Partition}:ecs:${Region}:${Account}:container-instance/${ClusterName}/${ContainerInstanceId}"},
		{Name: "ecs_daemon", Service: "ecs", Resource: "daemon", Template: "arn:${Partition}:ecs:${Region}:${Account}:daemon/${ClusterName}/${DaemonName}"},
		{Name: "ecs_daemon_deployment", Service: "ecs", Resource: "daemon-deployment", Template: "arn:${Partition}:ecs:${Region}:${Account}:daemon-deployment/${ClusterName}/${DaemonName}/${DaemonDeploymentId}"},
		{Name: "ecs_daemon_revision", Service: "ecs", Resource: "daemon-revision", Template: "arn:${Partition}:ecs:${Region}:${Account}:daemon-revision/${ClusterName}/${DaemonName}/${DaemonRevisionId}"},
		{Name: "ecs_daemon_task_definition", Service: "ecs", Resource: "daemon-task-definition", Template: "arn:${Partition}:ecs:${Region}:${Account}:daemon-task-definition/${DaemonTaskDefinitionFamilyName}:${DaemonTaskDefinitionRevisionNumber}"},
		{Name: "ecs_service", Service: "ecs", Resource: "service", Template: "arn:${Partition}:ecs:${Region}:${Account}:service/${ClusterName}/${ServiceName}"},
		{Name: "ecs_service_deployment", Service: "ecs", Resource: "service-deployment", Template: "arn:${Partition}:ecs:${Region}:${Account}:service-deployment/${ClusterName}/${ServiceName}/${ServiceDeploymentId}"},
		{Name: "ecs_service_revision", Service: "ecs", Resource: "service-revision", Template: "arn:${Partition}:ecs:${Region}:${Account}:service-revision/${ClusterName}/${ServiceName}/${ServiceRevisionId}"},
		{Name: "ecs_task", Service: "ecs", Resource: "task", Template: "arn:${Partition}:ecs:${Region}:${Account}:task/${ClusterName}/${TaskId}"},
		{Name: "ecs_task_definition", Service: "ecs", Resource: "task-definition", Template: "arn:${Partition}:ecs:${Region}:${Account}:task-definition/${TaskDefinitionFamilyName}:${TaskDefinitionRevisionNumber}"},
		{Name: "ecs_task_set", Service: "ecs", Resource: "task-set", Template: "arn:${Partition}:ecs:${Region}:${Account}:task-set/${ClusterName}/${ServiceName}/${TaskSetId}"},
	})
}
