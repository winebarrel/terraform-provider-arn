// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: partnercentral
// Source: https://servicereference.us-east-1.amazonaws.com/v1/partnercentral/partnercentral.json
// Functions: 23
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "partnercentral_benefit", Service: "partnercentral", Resource: "Benefit", Template: "arn:${Partition}:partnercentral:${Region}::catalog/${Catalog}/benefit/${Identifier}"},
		{Name: "partnercentral_benefit_allocation", Service: "partnercentral", Resource: "BenefitAllocation", Template: "arn:${Partition}:partnercentral:${Region}:${Account}:catalog/${Catalog}/benefit-allocation/${Identifier}"},
		{Name: "partnercentral_benefit_application", Service: "partnercentral", Resource: "BenefitApplication", Template: "arn:${Partition}:partnercentral:${Region}:${Account}:catalog/${Catalog}/benefit-application/${Identifier}"},
		{Name: "partnercentral_channel_handshake", Service: "partnercentral", Resource: "ChannelHandshake", Template: "arn:${Partition}:partnercentral:${Region}:${Account}:catalog/${Catalog}/channel-handshake/${Identifier}"},
		{Name: "partnercentral_connection", Service: "partnercentral", Resource: "Connection", Template: "arn:${Partition}:partnercentral:${Region}::catalog/${Catalog}/connection/${Identifier}"},
		{Name: "partnercentral_connection_invitation", Service: "partnercentral", Resource: "ConnectionInvitation", Template: "arn:${Partition}:partnercentral:${Region}::catalog/${Catalog}/connection-invitation/${Identifier}"},
		{Name: "partnercentral_connection_preferences", Service: "partnercentral", Resource: "ConnectionPreferences", Template: "arn:${Partition}:partnercentral:${Region}:${Account}:catalog/${Catalog}/connection-preferences"},
		{Name: "partnercentral_dashboard", Service: "partnercentral", Resource: "Dashboard", Template: "arn:${Partition}:partnercentral::${Account}:catalog/${Catalog}/ReportingData/${TableId}/Dashboard/${DashboardId}"},
		{Name: "partnercentral_engagement", Service: "partnercentral", Resource: "Engagement", Template: "arn:${Partition}:partnercentral:${Region}::catalog/${Catalog}/engagement/${Identifier}"},
		{Name: "partnercentral_engagement_by_accepting_invitation_task", Service: "partnercentral", Resource: "engagement-by-accepting-invitation-task", Template: "arn:${Partition}:partnercentral:${Region}::catalog/${Catalog}/engagement-by-accepting-invitation-task/${TaskId}"},
		{Name: "partnercentral_engagement_from_opportunity_task", Service: "partnercentral", Resource: "engagement-from-opportunity-task", Template: "arn:${Partition}:partnercentral:${Region}::catalog/${Catalog}/engagement-from-opportunity-task/${TaskId}"},
		{Name: "partnercentral_engagement_invitation", Service: "partnercentral", Resource: "engagement-invitation", Template: "arn:${Partition}:partnercentral:${Region}::catalog/${Catalog}/engagement-invitation/${Identifier}"},
		{Name: "partnercentral_marketplace_revenue_share", Service: "partnercentral", Resource: "MarketplaceRevenueShare", Template: "arn:${Partition}:partnercentral:${Region}:${Account}:catalog/${Catalog}/marketplace-revenue-share/${MarketplaceProductId}"},
		{Name: "partnercentral_opportunity", Service: "partnercentral", Resource: "Opportunity", Template: "arn:${Partition}:partnercentral:${Region}:${Account}:catalog/${Catalog}/opportunity/${Identifier}"},
		{Name: "partnercentral_opportunity_from_engagement_task", Service: "partnercentral", Resource: "OpportunityFromEngagementTask", Template: "arn:${Partition}:partnercentral:${Region}::catalog/${Catalog}/opportunity-from-engagement-task/${TaskId}"},
		{Name: "partnercentral_partner", Service: "partnercentral", Resource: "Partner", Template: "arn:${Partition}:partnercentral:${Region}:${Account}:catalog/${Catalog}/partner/${Identifier}"},
		{Name: "partnercentral_program_management_account", Service: "partnercentral", Resource: "ProgramManagementAccount", Template: "arn:${Partition}:partnercentral:${Region}:${Account}:catalog/${Catalog}/program-management-account/${Identifier}"},
		{Name: "partnercentral_prospecting_from_engagement_task", Service: "partnercentral", Resource: "ProspectingFromEngagementTask", Template: "arn:${Partition}:partnercentral:${Region}::catalog/${Catalog}/prospecting-from-engagement-task/${TaskIdentifier}"},
		{Name: "partnercentral_relationship", Service: "partnercentral", Resource: "Relationship", Template: "arn:${Partition}:partnercentral:${Region}:${Account}:catalog/${Catalog}/program-management-account/${ProgramManagementAccountId}/relationship/${RelationshipId}"},
		{Name: "partnercentral_resource_snapshot", Service: "partnercentral", Resource: "ResourceSnapshot", Template: "arn:${Partition}:partnercentral:${Region}:${Account}:catalog/${Catalog}/engagement/${EngagementIdentifier}/resource/${ResourceType}/${ResourceIdentifier}/template/${TemplateIdentifier}/resource-snapshot/${SnapshotRevision}"},
		{Name: "partnercentral_resource_snapshot_job", Service: "partnercentral", Resource: "resource-snapshot-job", Template: "arn:${Partition}:partnercentral:${Region}:${Account}:catalog/${Catalog}/resource-snapshot-job/${Identifier}"},
		{Name: "partnercentral_revenue_attribution", Service: "partnercentral", Resource: "RevenueAttribution", Template: "arn:${Partition}:partnercentral:${Region}:${Account}:catalog/${Catalog}/revenue-attribution/${RevenueAttributionId}"},
		{Name: "partnercentral_solution", Service: "partnercentral", Resource: "Solution", Template: "arn:${Partition}:partnercentral:${Region}:${Account}:catalog/${Catalog}/solution/${Identifier}"},
	})
}
