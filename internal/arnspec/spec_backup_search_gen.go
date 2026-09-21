// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: backup-search
// Source: https://servicereference.us-east-1.amazonaws.com/v1/backup-search/backup-search.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "backup_search_search_export_job", Service: "backup-search", Resource: "searchExportJob", Template: "arn:${Partition}:backup-search:${Region}:${Account}:search-export-job/${ResourceId}"},
		{Name: "backup_search_search_job", Service: "backup-search", Resource: "searchJob", Template: "arn:${Partition}:backup-search:${Region}:${Account}:search-job/${ResourceId}"},
	})
}
