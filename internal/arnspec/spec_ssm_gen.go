// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: ssm
// Source: https://servicereference.us-east-1.amazonaws.com/v1/ssm/ssm.json
// Functions: 22
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "ssm_association", Service: "ssm", Resource: "association", Template: "arn:${Partition}:ssm:${Region}:${Account}:association/${AssociationId}"},
		{Name: "ssm_automation_definition", Service: "ssm", Resource: "automation-definition", Template: "arn:${Partition}:ssm:${Region}:${Account}:automation-definition/${AutomationDefinitionName}:${VersionId}"},
		{Name: "ssm_automation_execution", Service: "ssm", Resource: "automation-execution", Template: "arn:${Partition}:ssm:${Region}:${Account}:automation-execution/${AutomationExecutionId}"},
		{Name: "ssm_bucket", Service: "ssm", Resource: "bucket", Template: "arn:${Partition}:s3:::${BucketName}"},
		{Name: "ssm_cloud_connector", Service: "ssm", Resource: "cloud-connector", Template: "arn:${Partition}:ssm:${Region}:${Account}:cloud-connector/${CloudConnectorId}"},
		{Name: "ssm_document", Service: "ssm", Resource: "document", Template: "arn:${Partition}:ssm:${Region}:${Account}:document/${DocumentName}"},
		{Name: "ssm_iam_role", Service: "ssm", Resource: "iam-role", Template: "arn:${Partition}:iam::${Account}:role/${RoleName}"},
		{Name: "ssm_instance", Service: "ssm", Resource: "instance", Template: "arn:${Partition}:ec2:${Region}:${Account}:instance/${InstanceId}"},
		{Name: "ssm_maintenancewindow", Service: "ssm", Resource: "maintenancewindow", Template: "arn:${Partition}:ssm:${Region}:${Account}:maintenancewindow/${ResourceId}"},
		{Name: "ssm_managed_instance", Service: "ssm", Resource: "managed-instance", Template: "arn:${Partition}:ssm:${Region}:${Account}:managed-instance/${InstanceId}"},
		{Name: "ssm_managed_instance_inventory", Service: "ssm", Resource: "managed-instance-inventory", Template: "arn:${Partition}:ssm:${Region}:${Account}:managed-instance-inventory/${InstanceId}"},
		{Name: "ssm_opsitem", Service: "ssm", Resource: "opsitem", Template: "arn:${Partition}:ssm:${Region}:${Account}:opsitem/${ResourceId}"},
		{Name: "ssm_opsitemgroup", Service: "ssm", Resource: "opsitemgroup", Template: "arn:${Partition}:ssm:${Region}:${Account}:opsitemgroup/default"},
		{Name: "ssm_opsmetadata", Service: "ssm", Resource: "opsmetadata", Template: "arn:${Partition}:ssm:${Region}:${Account}:opsmetadata/${ResourceId}"},
		{Name: "ssm_parameter", Service: "ssm", Resource: "parameter", Template: "arn:${Partition}:ssm:${Region}:${Account}:parameter/${ParameterNameWithoutLeadingSlash}"},
		{Name: "ssm_patchbaseline", Service: "ssm", Resource: "patchbaseline", Template: "arn:${Partition}:ssm:${Region}:${Account}:patchbaseline/${PatchBaselineIdResourceId}"},
		{Name: "ssm_resourcedatasync", Service: "ssm", Resource: "resourcedatasync", Template: "arn:${Partition}:ssm:${Region}:${Account}:resource-data-sync/${SyncName}"},
		{Name: "ssm_servicesetting", Service: "ssm", Resource: "servicesetting", Template: "arn:${Partition}:ssm:${Region}:${Account}:servicesetting/${ResourceId}"},
		{Name: "ssm_session", Service: "ssm", Resource: "session", Template: "arn:${Partition}:ssm:${Region}:${Account}:session/${SessionId}"},
		{Name: "ssm_task", Service: "ssm", Resource: "task", Template: "arn:${Partition}:ecs:${Region}:${Account}:task/${TaskId}"},
		{Name: "ssm_windowtarget", Service: "ssm", Resource: "windowtarget", Template: "arn:${Partition}:ssm:${Region}:${Account}:windowtarget/${WindowTargetId}"},
		{Name: "ssm_windowtask", Service: "ssm", Resource: "windowtask", Template: "arn:${Partition}:ssm:${Region}:${Account}:windowtask/${WindowTaskId}"},
	})
}
