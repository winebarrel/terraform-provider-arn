// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: outposts
// Source: https://servicereference.us-east-1.amazonaws.com/v1/outposts/outposts.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "outposts_outpost", Service: "outposts", Resource: "outpost", Template: "arn:${Partition}:outposts:${Region}:${Account}:outpost/${OutpostId}"},
		{Name: "outposts_site", Service: "outposts", Resource: "site", Template: "arn:${Partition}:outposts:${Region}:${Account}:site/${SiteId}"},
	})
}
