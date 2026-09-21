// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: codeguru-profiler
// Source: https://servicereference.us-east-1.amazonaws.com/v1/codeguru-profiler/codeguru-profiler.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "codeguru_profiler_profiling_group", Service: "codeguru-profiler", Resource: "ProfilingGroup", Template: "arn:${Partition}:codeguru-profiler:${Region}:${Account}:profilingGroup/${ProfilingGroupName}"},
	})
}
