// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: autoscaling
// Source: https://servicereference.us-east-1.amazonaws.com/v1/autoscaling/autoscaling.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "autoscaling_auto_scaling_group", Service: "autoscaling", Resource: "autoScalingGroup", Template: "arn:${Partition}:autoscaling:${Region}:${Account}:autoScalingGroup:${GroupId}:autoScalingGroupName/${GroupFriendlyName}"},
		{Name: "autoscaling_launch_configuration", Service: "autoscaling", Resource: "launchConfiguration", Template: "arn:${Partition}:autoscaling:${Region}:${Account}:launchConfiguration:${Id}:launchConfigurationName/${LaunchConfigurationName}"},
	})
}
