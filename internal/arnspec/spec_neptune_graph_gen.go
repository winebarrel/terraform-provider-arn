// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: neptune-graph
// Source: https://servicereference.us-east-1.amazonaws.com/v1/neptune-graph/neptune-graph.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "neptune_graph_export_task", Service: "neptune-graph", Resource: "export-task", Template: "arn:${Partition}:neptune-graph:${Region}:${Account}:export-task/${ResourceId}"},
		{Name: "neptune_graph_graph", Service: "neptune-graph", Resource: "graph", Template: "arn:${Partition}:neptune-graph:${Region}:${Account}:graph/${ResourceId}"},
		{Name: "neptune_graph_graph_snapshot", Service: "neptune-graph", Resource: "graph-snapshot", Template: "arn:${Partition}:neptune-graph:${Region}:${Account}:graph-snapshot/${ResourceId}"},
		{Name: "neptune_graph_import_task", Service: "neptune-graph", Resource: "import-task", Template: "arn:${Partition}:neptune-graph:${Region}:${Account}:import-task/${ResourceId}"},
	})
}
