// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: appstream
// Source: https://servicereference.us-east-1.amazonaws.com/v1/appstream/appstream.json
// Functions: 7
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "appstream_app_block", Service: "appstream", Resource: "app-block", Template: "arn:${Partition}:appstream:${Region}:${Account}:app-block/${AppBlockName}"},
		{Name: "appstream_app_block_builder", Service: "appstream", Resource: "app-block-builder", Template: "arn:${Partition}:appstream:${Region}:${Account}:app-block-builder/${AppBlockBuilderName}"},
		{Name: "appstream_application", Service: "appstream", Resource: "application", Template: "arn:${Partition}:appstream:${Region}:${Account}:application/${ApplicationName}"},
		{Name: "appstream_fleet", Service: "appstream", Resource: "fleet", Template: "arn:${Partition}:appstream:${Region}:${Account}:fleet/${FleetName}"},
		{Name: "appstream_image", Service: "appstream", Resource: "image", Template: "arn:${Partition}:appstream:${Region}:${Account}:image/${ImageName}"},
		{Name: "appstream_image_builder", Service: "appstream", Resource: "image-builder", Template: "arn:${Partition}:appstream:${Region}:${Account}:image-builder/${ImageBuilderName}"},
		{Name: "appstream_stack", Service: "appstream", Resource: "stack", Template: "arn:${Partition}:appstream:${Region}:${Account}:stack/${StackName}"},
	})
}
