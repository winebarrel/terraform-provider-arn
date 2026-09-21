// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: simspaceweaver
// Source: https://servicereference.us-east-1.amazonaws.com/v1/simspaceweaver/simspaceweaver.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "simspaceweaver_simulation", Service: "simspaceweaver", Resource: "Simulation", Template: "arn:${Partition}:simspaceweaver:${Region}:${Account}:simulation/${SimulationName}"},
	})
}
