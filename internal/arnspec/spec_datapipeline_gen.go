// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: datapipeline
// Source: https://servicereference.us-east-1.amazonaws.com/v1/datapipeline/datapipeline.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "datapipeline_pipeline", Service: "datapipeline", Resource: "pipeline", Template: "arn:${Partition}:datapipeline:${Region}:${Account}:pipeline/${PipelineId}"},
	})
}
