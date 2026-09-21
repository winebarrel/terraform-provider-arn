// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: codebuild
// Source: https://servicereference.us-east-1.amazonaws.com/v1/codebuild/codebuild.json
// Functions: 7
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "codebuild_build", Service: "codebuild", Resource: "build", Template: "arn:${Partition}:codebuild:${Region}:${Account}:build/${BuildId}"},
		{Name: "codebuild_build_batch", Service: "codebuild", Resource: "build-batch", Template: "arn:${Partition}:codebuild:${Region}:${Account}:build-batch/${BuildBatchId}"},
		{Name: "codebuild_fleet", Service: "codebuild", Resource: "fleet", Template: "arn:${Partition}:codebuild:${Region}:${Account}:fleet/${FleetName}:${FleetId}"},
		{Name: "codebuild_project", Service: "codebuild", Resource: "project", Template: "arn:${Partition}:codebuild:${Region}:${Account}:project/${ProjectName}"},
		{Name: "codebuild_report", Service: "codebuild", Resource: "report", Template: "arn:${Partition}:codebuild:${Region}:${Account}:report/${ReportGroupName}:${ReportId}"},
		{Name: "codebuild_report_group", Service: "codebuild", Resource: "report-group", Template: "arn:${Partition}:codebuild:${Region}:${Account}:report-group/${ReportGroupName}"},
		{Name: "codebuild_sandbox", Service: "codebuild", Resource: "sandbox", Template: "arn:${Partition}:codebuild:${Region}:${Account}:sandbox/${SandboxId}"},
	})
}
