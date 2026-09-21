// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: finspace
// Source: https://servicereference.us-east-1.amazonaws.com/v1/finspace/finspace.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "finspace_environment", Service: "finspace", Resource: "environment", Template: "arn:${Partition}:finspace:${Region}:${Account}:environment/${EnvironmentId}"},
		{Name: "finspace_kx_cluster", Service: "finspace", Resource: "kxCluster", Template: "arn:${Partition}:finspace:${Region}:${Account}:kxEnvironment/${EnvironmentId}/kxCluster/${KxCluster}"},
		{Name: "finspace_kx_database", Service: "finspace", Resource: "kxDatabase", Template: "arn:${Partition}:finspace:${Region}:${Account}:kxEnvironment/${EnvironmentId}/kxDatabase/${KxDatabase}"},
		{Name: "finspace_kx_dataview", Service: "finspace", Resource: "kxDataview", Template: "arn:${Partition}:finspace:${Region}:${Account}:kxEnvironment/${EnvironmentId}/kxDatabase/${KxDatabase}/kxDataview/${KxDataview}"},
		{Name: "finspace_kx_environment", Service: "finspace", Resource: "kxEnvironment", Template: "arn:${Partition}:finspace:${Region}:${Account}:kxEnvironment/${EnvironmentId}"},
		{Name: "finspace_kx_scaling_group", Service: "finspace", Resource: "kxScalingGroup", Template: "arn:${Partition}:finspace:${Region}:${Account}:kxEnvironment/${EnvironmentId}/kxScalingGroup/${KxScalingGroup}"},
		{Name: "finspace_kx_user", Service: "finspace", Resource: "kxUser", Template: "arn:${Partition}:finspace:${Region}:${Account}:kxEnvironment/${EnvironmentId}/kxUser/${UserName}"},
		{Name: "finspace_kx_volume", Service: "finspace", Resource: "kxVolume", Template: "arn:${Partition}:finspace:${Region}:${Account}:kxEnvironment/${EnvironmentId}/kxVolume/${KxVolume}"},
		{Name: "finspace_user", Service: "finspace", Resource: "user", Template: "arn:${Partition}:finspace:${Region}:${Account}:user/${UserId}"},
	})
}
