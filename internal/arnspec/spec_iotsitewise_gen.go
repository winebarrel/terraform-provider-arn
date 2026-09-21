// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: iotsitewise
// Source: https://servicereference.us-east-1.amazonaws.com/v1/iotsitewise/iotsitewise.json
// Functions: 14
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "iotsitewise_access_policy", Service: "iotsitewise", Resource: "access-policy", Template: "arn:${Partition}:iotsitewise:${Region}:${Account}:access-policy/${AccessPolicyId}"},
		{Name: "iotsitewise_application", Service: "iotsitewise", Resource: "application", Template: "arn:${Partition}:iotsitewise:${Region}:${Account}:workspace/${WorkspaceName}/application/${ApplicationId}"},
		{Name: "iotsitewise_asset", Service: "iotsitewise", Resource: "asset", Template: "arn:${Partition}:iotsitewise:${Region}:${Account}:asset/${AssetId}"},
		{Name: "iotsitewise_asset_model", Service: "iotsitewise", Resource: "asset-model", Template: "arn:${Partition}:iotsitewise:${Region}:${Account}:asset-model/${AssetModelId}"},
		{Name: "iotsitewise_computation_model", Service: "iotsitewise", Resource: "computation-model", Template: "arn:${Partition}:iotsitewise:${Region}:${Account}:computation-model/${ComputationModelId}"},
		{Name: "iotsitewise_dashboard", Service: "iotsitewise", Resource: "dashboard", Template: "arn:${Partition}:iotsitewise:${Region}:${Account}:dashboard/${DashboardId}"},
		{Name: "iotsitewise_dataset", Service: "iotsitewise", Resource: "dataset", Template: "arn:${Partition}:iotsitewise:${Region}:${Account}:dataset/${DatasetId}"},
		{Name: "iotsitewise_gateway", Service: "iotsitewise", Resource: "gateway", Template: "arn:${Partition}:iotsitewise:${Region}:${Account}:gateway/${GatewayId}"},
		{Name: "iotsitewise_pipeline", Service: "iotsitewise", Resource: "pipeline", Template: "arn:${Partition}:iotsitewise:${Region}:${Account}:workspace/${WorkspaceName}/pipeline/${PipelineName}"},
		{Name: "iotsitewise_portal", Service: "iotsitewise", Resource: "portal", Template: "arn:${Partition}:iotsitewise:${Region}:${Account}:portal/${PortalId}"},
		{Name: "iotsitewise_project", Service: "iotsitewise", Resource: "project", Template: "arn:${Partition}:iotsitewise:${Region}:${Account}:project/${ProjectId}"},
		{Name: "iotsitewise_task", Service: "iotsitewise", Resource: "task", Template: "arn:${Partition}:iotsitewise:${Region}:${Account}:workspace/${WorkspaceName}/task/${TaskName}"},
		{Name: "iotsitewise_time_series", Service: "iotsitewise", Resource: "time-series", Template: "arn:${Partition}:iotsitewise:${Region}:${Account}:time-series/${TimeSeriesId}"},
		{Name: "iotsitewise_workspace", Service: "iotsitewise", Resource: "workspace", Template: "arn:${Partition}:iotsitewise:${Region}:${Account}:workspace/${WorkspaceName}"},
	})
}
