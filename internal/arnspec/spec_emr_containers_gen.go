// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: emr-containers
// Source: https://servicereference.us-east-1.amazonaws.com/v1/emr-containers/emr-containers.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "emr_containers_job_run", Service: "emr-containers", Resource: "jobRun", Template: "arn:${Partition}:emr-containers:${Region}:${Account}:/virtualclusters/${VirtualClusterId}/jobruns/${JobRunId}"},
		{Name: "emr_containers_job_template", Service: "emr-containers", Resource: "jobTemplate", Template: "arn:${Partition}:emr-containers:${Region}:${Account}:/jobtemplates/${JobTemplateId}"},
		{Name: "emr_containers_managed_endpoint", Service: "emr-containers", Resource: "managedEndpoint", Template: "arn:${Partition}:emr-containers:${Region}:${Account}:/virtualclusters/${VirtualClusterId}/endpoints/${EndpointId}"},
		{Name: "emr_containers_security_configuration", Service: "emr-containers", Resource: "securityConfiguration", Template: "arn:${Partition}:emr-containers:${Region}:${Account}:/securityconfigurations/${SecurityConfigurationId}"},
		{Name: "emr_containers_virtual_cluster", Service: "emr-containers", Resource: "virtualCluster", Template: "arn:${Partition}:emr-containers:${Region}:${Account}:/virtualclusters/${VirtualClusterId}"},
	})
}
