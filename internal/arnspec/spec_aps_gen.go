// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: aps
// Source: https://servicereference.us-east-1.amazonaws.com/v1/aps/aps.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "aps_anomalydetector", Service: "aps", Resource: "anomalydetector", Template: "arn:${Partition}:aps:${Region}:${Account}:anomalydetector/${WorkspaceId}/${AnomalyDetectorId}"},
		{Name: "aps_cluster", Service: "aps", Resource: "cluster", Template: "arn:${Partition}:eks:${Region}:${Account}:cluster/${ClusterName}"},
		{Name: "aps_rulegroupsnamespace", Service: "aps", Resource: "rulegroupsnamespace", Template: "arn:${Partition}:aps:${Region}:${Account}:rulegroupsnamespace/${WorkspaceId}/${Namespace}"},
		{Name: "aps_scraper", Service: "aps", Resource: "scraper", Template: "arn:${Partition}:aps:${Region}:${Account}:scraper/${ScraperId}"},
		{Name: "aps_workspace", Service: "aps", Resource: "workspace", Template: "arn:${Partition}:aps:${Region}:${Account}:workspace/${WorkspaceId}"},
	})
}
