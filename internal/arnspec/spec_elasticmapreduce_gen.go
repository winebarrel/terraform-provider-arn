// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: elasticmapreduce
// Source: https://servicereference.us-east-1.amazonaws.com/v1/elasticmapreduce/elasticmapreduce.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "elasticmapreduce_cluster", Service: "elasticmapreduce", Resource: "cluster", Template: "arn:${Partition}:elasticmapreduce:${Region}:${Account}:cluster/${ClusterId}"},
		{Name: "elasticmapreduce_editor", Service: "elasticmapreduce", Resource: "editor", Template: "arn:${Partition}:elasticmapreduce:${Region}:${Account}:editor/${EditorId}"},
		{Name: "elasticmapreduce_notebook_execution", Service: "elasticmapreduce", Resource: "notebook-execution", Template: "arn:${Partition}:elasticmapreduce:${Region}:${Account}:notebook-execution/${NotebookExecutionId}"},
		{Name: "elasticmapreduce_session", Service: "elasticmapreduce", Resource: "session", Template: "arn:${Partition}:elasticmapreduce:${Region}:${Account}:cluster/${ClusterId}/session/${SessionId}"},
		{Name: "elasticmapreduce_studio", Service: "elasticmapreduce", Resource: "studio", Template: "arn:${Partition}:elasticmapreduce:${Region}:${Account}:studio/${StudioId}"},
	})
}
