// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: macie2
// Source: https://servicereference.us-east-1.amazonaws.com/v1/macie2/macie2.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "macie2_allow_list", Service: "macie2", Resource: "AllowList", Template: "arn:${Partition}:macie2:${Region}:${Account}:allow-list/${ResourceId}"},
		{Name: "macie2_classification_job", Service: "macie2", Resource: "ClassificationJob", Template: "arn:${Partition}:macie2:${Region}:${Account}:classification-job/${ResourceId}"},
		{Name: "macie2_custom_data_identifier", Service: "macie2", Resource: "CustomDataIdentifier", Template: "arn:${Partition}:macie2:${Region}:${Account}:custom-data-identifier/${ResourceId}"},
		{Name: "macie2_findings_filter", Service: "macie2", Resource: "FindingsFilter", Template: "arn:${Partition}:macie2:${Region}:${Account}:findings-filter/${ResourceId}"},
		{Name: "macie2_member", Service: "macie2", Resource: "Member", Template: "arn:${Partition}:macie2:${Region}:${Account}:member/${ResourceId}"},
	})
}
