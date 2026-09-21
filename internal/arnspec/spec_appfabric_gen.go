// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: appfabric
// Source: https://servicereference.us-east-1.amazonaws.com/v1/appfabric/appfabric.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "appfabric_appauthorization", Service: "appfabric", Resource: "appauthorization", Template: "arn:${Partition}:appfabric:${Region}:${Account}:appbundle/${AppbundleId}/appauthorization/${AppAuthorizationIdentifier}"},
		{Name: "appfabric_appbundle", Service: "appfabric", Resource: "appbundle", Template: "arn:${Partition}:appfabric:${Region}:${Account}:appbundle/${AppBundleIdentifier}"},
		{Name: "appfabric_ingestion", Service: "appfabric", Resource: "ingestion", Template: "arn:${Partition}:appfabric:${Region}:${Account}:appbundle/${AppbundleId}/ingestion/${IngestionIdentifier}"},
		{Name: "appfabric_ingestiondestination", Service: "appfabric", Resource: "ingestiondestination", Template: "arn:${Partition}:appfabric:${Region}:${Account}:appbundle/${AppbundleId}/ingestion/${IngestionIdentifier}/ingestiondestination/${IngestionDestinationIdentifier}"},
	})
}
