// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: robomaker
// Source: https://servicereference.us-east-1.amazonaws.com/v1/robomaker/robomaker.json
// Functions: 11
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "robomaker_deployment_fleet", Service: "robomaker", Resource: "deploymentFleet", Template: "arn:${Partition}:robomaker:${Region}:${Account}:deployment-fleet/${FleetName}/${CreatedOnEpoch}"},
		{Name: "robomaker_deployment_job", Service: "robomaker", Resource: "deploymentJob", Template: "arn:${Partition}:robomaker:${Region}:${Account}:deployment-job/${DeploymentJobId}"},
		{Name: "robomaker_robot", Service: "robomaker", Resource: "robot", Template: "arn:${Partition}:robomaker:${Region}:${Account}:robot/${RobotName}/${CreatedOnEpoch}"},
		{Name: "robomaker_robot_application", Service: "robomaker", Resource: "robotApplication", Template: "arn:${Partition}:robomaker:${Region}:${Account}:robot-application/${ApplicationName}/${CreatedOnEpoch}"},
		{Name: "robomaker_simulation_application", Service: "robomaker", Resource: "simulationApplication", Template: "arn:${Partition}:robomaker:${Region}:${Account}:simulation-application/${ApplicationName}/${CreatedOnEpoch}"},
		{Name: "robomaker_simulation_job", Service: "robomaker", Resource: "simulationJob", Template: "arn:${Partition}:robomaker:${Region}:${Account}:simulation-job/${SimulationJobId}"},
		{Name: "robomaker_simulation_job_batch", Service: "robomaker", Resource: "simulationJobBatch", Template: "arn:${Partition}:robomaker:${Region}:${Account}:simulation-job-batch/${SimulationJobBatchId}"},
		{Name: "robomaker_world", Service: "robomaker", Resource: "world", Template: "arn:${Partition}:robomaker:${Region}:${Account}:world/${WorldId}"},
		{Name: "robomaker_world_export_job", Service: "robomaker", Resource: "worldExportJob", Template: "arn:${Partition}:robomaker:${Region}:${Account}:world-export-job/${WorldExportJobId}"},
		{Name: "robomaker_world_generation_job", Service: "robomaker", Resource: "worldGenerationJob", Template: "arn:${Partition}:robomaker:${Region}:${Account}:world-generation-job/${WorldGenerationJobId}"},
		{Name: "robomaker_world_template", Service: "robomaker", Resource: "worldTemplate", Template: "arn:${Partition}:robomaker:${Region}:${Account}:world-template/${WorldTemplateJobId}"},
	})
}
