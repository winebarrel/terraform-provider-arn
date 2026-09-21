// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: codepipeline
// Source: https://servicereference.us-east-1.amazonaws.com/v1/codepipeline/codepipeline.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "codepipeline_action", Service: "codepipeline", Resource: "action", Template: "arn:${Partition}:codepipeline:${Region}:${Account}:${PipelineName}/${StageName}/${ActionName}"},
		{Name: "codepipeline_actiontype", Service: "codepipeline", Resource: "actiontype", Template: "arn:${Partition}:codepipeline:${Region}:${Account}:actiontype:${Owner}/${Category}/${Provider}/${Version}"},
		{Name: "codepipeline_pipeline", Service: "codepipeline", Resource: "pipeline", Template: "arn:${Partition}:codepipeline:${Region}:${Account}:${PipelineName}"},
		{Name: "codepipeline_stage", Service: "codepipeline", Resource: "stage", Template: "arn:${Partition}:codepipeline:${Region}:${Account}:${PipelineName}/${StageName}"},
		{Name: "codepipeline_webhook", Service: "codepipeline", Resource: "webhook", Template: "arn:${Partition}:codepipeline:${Region}:${Account}:webhook:${WebhookName}"},
	})
}
