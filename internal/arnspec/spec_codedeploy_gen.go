// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: codedeploy
// Source: https://servicereference.us-east-1.amazonaws.com/v1/codedeploy/codedeploy.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "codedeploy_application", Service: "codedeploy", Resource: "application", Template: "arn:${Partition}:codedeploy:${Region}:${Account}:application:${ApplicationName}"},
		{Name: "codedeploy_deploymentconfig", Service: "codedeploy", Resource: "deploymentconfig", Template: "arn:${Partition}:codedeploy:${Region}:${Account}:deploymentconfig:${DeploymentConfigurationName}"},
		{Name: "codedeploy_deploymentgroup", Service: "codedeploy", Resource: "deploymentgroup", Template: "arn:${Partition}:codedeploy:${Region}:${Account}:deploymentgroup:${ApplicationName}/${DeploymentGroupName}"},
		{Name: "codedeploy_instance", Service: "codedeploy", Resource: "instance", Template: "arn:${Partition}:codedeploy:${Region}:${Account}:instance:${InstanceName}"},
	})
}
