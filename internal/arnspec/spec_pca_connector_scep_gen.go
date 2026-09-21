// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: pca-connector-scep
// Source: https://servicereference.us-east-1.amazonaws.com/v1/pca-connector-scep/pca-connector-scep.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "pca_connector_scep_challenge", Service: "pca-connector-scep", Resource: "Challenge", Template: "arn:${Partition}:pca-connector-scep:${Region}:${Account}:connector/${ConnectorId}/challenge/${ChallengeId}"},
		{Name: "pca_connector_scep_connector", Service: "pca-connector-scep", Resource: "Connector", Template: "arn:${Partition}:pca-connector-scep:${Region}:${Account}:connector/${ConnectorId}"},
	})
}
