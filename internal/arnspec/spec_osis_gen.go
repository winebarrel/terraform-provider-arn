// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: osis
// Source: https://servicereference.us-east-1.amazonaws.com/v1/osis/osis.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "osis_pipeline", Service: "osis", Resource: "pipeline", Template: "arn:${Partition}:osis:${Region}:${Account}:pipeline/${PipelineName}"},
		{Name: "osis_pipeline_blueprint", Service: "osis", Resource: "pipeline-blueprint", Template: "arn:${Partition}:osis:${Region}:${Account}:blueprint/${BlueprintName}"},
		{Name: "osis_pipeline_endpoint", Service: "osis", Resource: "pipeline-endpoint", Template: "arn:${Partition}:osis:${Region}:${Account}:endpoint/${EndpointId}"},
	})
}
