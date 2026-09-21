// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: access-analyzer
// Source: https://servicereference.us-east-1.amazonaws.com/v1/access-analyzer/access-analyzer.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "access_analyzer_analyzer", Service: "access-analyzer", Resource: "Analyzer", Template: "arn:${Partition}:access-analyzer:${Region}:${Account}:analyzer/${AnalyzerName}"},
		{Name: "access_analyzer_archive_rule", Service: "access-analyzer", Resource: "ArchiveRule", Template: "arn:${Partition}:access-analyzer:${Region}:${Account}:analyzer/${AnalyzerName}/archive-rule/${RuleName}"},
	})
}
