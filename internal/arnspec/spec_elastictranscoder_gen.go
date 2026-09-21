// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: elastictranscoder
// Source: https://servicereference.us-east-1.amazonaws.com/v1/elastictranscoder/elastictranscoder.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "elastictranscoder_job", Service: "elastictranscoder", Resource: "job", Template: "arn:${Partition}:elastictranscoder:${Region}:${Account}:job/${JobId}"},
		{Name: "elastictranscoder_pipeline", Service: "elastictranscoder", Resource: "pipeline", Template: "arn:${Partition}:elastictranscoder:${Region}:${Account}:pipeline/${PipelineId}"},
		{Name: "elastictranscoder_preset", Service: "elastictranscoder", Resource: "preset", Template: "arn:${Partition}:elastictranscoder:${Region}:${Account}:preset/${PresetId}"},
	})
}
