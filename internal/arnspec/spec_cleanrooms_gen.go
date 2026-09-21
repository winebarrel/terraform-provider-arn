// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cleanrooms
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cleanrooms/cleanrooms.json
// Functions: 10
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cleanrooms_analysistemplate", Service: "cleanrooms", Resource: "analysistemplate", Template: "arn:${Partition}:cleanrooms:${Region}:${Account}:membership/${MembershipId}/analysistemplate/${AnalysisTemplateId}"},
		{Name: "cleanrooms_collaboration", Service: "cleanrooms", Resource: "collaboration", Template: "arn:${Partition}:cleanrooms:${Region}:${Account}:collaboration/${CollaborationId}"},
		{Name: "cleanrooms_configuredaudiencemodelassociation", Service: "cleanrooms", Resource: "configuredaudiencemodelassociation", Template: "arn:${Partition}:cleanrooms:${Region}:${Account}:membership/${MembershipId}/configuredaudiencemodelassociation/${ConfiguredAudienceModelAssociationId}"},
		{Name: "cleanrooms_configuredtable", Service: "cleanrooms", Resource: "configuredtable", Template: "arn:${Partition}:cleanrooms:${Region}:${Account}:configuredtable/${ConfiguredTableId}"},
		{Name: "cleanrooms_configuredtableassociation", Service: "cleanrooms", Resource: "configuredtableassociation", Template: "arn:${Partition}:cleanrooms:${Region}:${Account}:membership/${MembershipId}/configuredtableassociation/${ConfiguredTableAssociationId}"},
		{Name: "cleanrooms_idmappingtable", Service: "cleanrooms", Resource: "idmappingtable", Template: "arn:${Partition}:cleanrooms:${Region}:${Account}:membership/${MembershipId}/idmappingtable/${IdMappingTableId}"},
		{Name: "cleanrooms_idnamespaceassociation", Service: "cleanrooms", Resource: "idnamespaceassociation", Template: "arn:${Partition}:cleanrooms:${Region}:${Account}:membership/${MembershipId}/idnamespaceassociation/${IdNamespaceAssociationId}"},
		{Name: "cleanrooms_intermediatetable", Service: "cleanrooms", Resource: "intermediatetable", Template: "arn:${Partition}:cleanrooms:${Region}:${Account}:membership/${MembershipId}/intermediatetable/${IntermediateTableId}"},
		{Name: "cleanrooms_membership", Service: "cleanrooms", Resource: "membership", Template: "arn:${Partition}:cleanrooms:${Region}:${Account}:membership/${MembershipId}"},
		{Name: "cleanrooms_privacybudgettemplate", Service: "cleanrooms", Resource: "privacybudgettemplate", Template: "arn:${Partition}:cleanrooms:${Region}:${Account}:membership/${MembershipId}/privacybudgettemplate/${PrivacyBudgetTemplateId}"},
	})
}
