// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: pca-connector-ad
// Source: https://servicereference.us-east-1.amazonaws.com/v1/pca-connector-ad/pca-connector-ad.json
// Functions: 3
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "pca_connector_ad_connector", Service: "pca-connector-ad", Resource: "Connector", Template: "arn:${Partition}:pca-connector-ad:${Region}:${Account}:connector/${ConnectorId}"},
		{Name: "pca_connector_ad_directory_registration", Service: "pca-connector-ad", Resource: "DirectoryRegistration", Template: "arn:${Partition}:pca-connector-ad:${Region}:${Account}:directory-registration/${DirectoryId}"},
		{Name: "pca_connector_ad_template", Service: "pca-connector-ad", Resource: "Template", Template: "arn:${Partition}:pca-connector-ad:${Region}:${Account}:connector/${ConnectorId}/template/${TemplateId}"},
	})
}
