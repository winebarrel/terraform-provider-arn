// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: grafana
// Source: https://servicereference.us-east-1.amazonaws.com/v1/grafana/grafana.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "grafana_workspace", Service: "grafana", Resource: "workspace", Template: "arn:${Partition}:grafana:${Region}:${Account}:/workspaces/${ResourceId}"},
	})
}
