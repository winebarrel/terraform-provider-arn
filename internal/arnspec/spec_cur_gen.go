// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cur
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cur/cur.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cur_cur", Service: "cur", Resource: "cur", Template: "arn:${Partition}:cur:${Region}:${Account}:definition/${ReportName}"},
	})
}
