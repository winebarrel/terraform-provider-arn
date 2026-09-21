// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: imagebuilder
// Source: https://servicereference.us-east-1.amazonaws.com/v1/imagebuilder/imagebuilder.json
// Functions: 16
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "imagebuilder_all_component_build_versions", Service: "imagebuilder", Resource: "allComponentBuildVersions", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:component/${ComponentName}/${ComponentVersion}/*"},
		{Name: "imagebuilder_all_image_build_versions", Service: "imagebuilder", Resource: "allImageBuildVersions", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:image/${ImageName}/${ImageVersion}/*"},
		{Name: "imagebuilder_all_workflow_build_versions", Service: "imagebuilder", Resource: "allWorkflowBuildVersions", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:workflow/${WorkflowType}/${WorkflowName}/${WorkflowVersion}/*"},
		{Name: "imagebuilder_component", Service: "imagebuilder", Resource: "component", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:component/${ComponentName}/${ComponentVersion}/${ComponentBuildVersion}"},
		{Name: "imagebuilder_container_recipe", Service: "imagebuilder", Resource: "containerRecipe", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:container-recipe/${ContainerRecipeName}/${ContainerRecipeVersion}"},
		{Name: "imagebuilder_distribution_configuration", Service: "imagebuilder", Resource: "distributionConfiguration", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:distribution-configuration/${DistributionConfigurationName}"},
		{Name: "imagebuilder_image", Service: "imagebuilder", Resource: "image", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:image/${ImageName}/${ImageVersion}/${ImageBuildVersion}"},
		{Name: "imagebuilder_image_pipeline", Service: "imagebuilder", Resource: "imagePipeline", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:image-pipeline/${ImagePipelineName}"},
		{Name: "imagebuilder_image_recipe", Service: "imagebuilder", Resource: "imageRecipe", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:image-recipe/${ImageRecipeName}/${ImageRecipeVersion}"},
		{Name: "imagebuilder_image_version", Service: "imagebuilder", Resource: "imageVersion", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:image/${ImageName}/${ImageVersion}"},
		{Name: "imagebuilder_infrastructure_configuration", Service: "imagebuilder", Resource: "infrastructureConfiguration", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:infrastructure-configuration/${ResourceId}"},
		{Name: "imagebuilder_lifecycle_execution", Service: "imagebuilder", Resource: "lifecycleExecution", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:lifecycle-execution/${LifecycleExecutionId}"},
		{Name: "imagebuilder_lifecycle_policy", Service: "imagebuilder", Resource: "lifecyclePolicy", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:lifecycle-policy/${LifecyclePolicyName}"},
		{Name: "imagebuilder_workflow", Service: "imagebuilder", Resource: "workflow", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:workflow/${WorkflowType}/${WorkflowName}/${WorkflowVersion}/${WorkflowBuildVersion}"},
		{Name: "imagebuilder_workflow_execution", Service: "imagebuilder", Resource: "workflowExecution", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:workflow-execution/${WorkflowExecutionId}"},
		{Name: "imagebuilder_workflow_step_execution", Service: "imagebuilder", Resource: "workflowStepExecution", Template: "arn:${Partition}:imagebuilder:${Region}:${Account}:workflow-step-execution/${WorkflowStepExecutionId}"},
	})
}
