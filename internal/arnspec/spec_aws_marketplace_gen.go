// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: aws-marketplace
// Source: https://servicereference.us-east-1.amazonaws.com/v1/aws-marketplace/aws-marketplace.json
// Functions: 18
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "aws_marketplace_all_listings", Service: "aws-marketplace", Resource: "AllListings", Template: "arn:${Partition}:aws-marketplace:::catalog/${CatalogName}/listing/*"},
		{Name: "aws_marketplace_all_purchase_options", Service: "aws-marketplace", Resource: "AllPurchaseOptions", Template: "arn:${Partition}:aws-marketplace:::catalog/${CatalogName}/purchaseOption/*"},
		{Name: "aws_marketplace_assessment", Service: "aws-marketplace", Resource: "Assessment", Template: "arn:${Partition}:aws-marketplace:${Region}::${Catalog}/Assessment/${ResourceId}"},
		{Name: "aws_marketplace_change_set", Service: "aws-marketplace", Resource: "ChangeSet", Template: "arn:${Partition}:aws-marketplace:${Region}:${Account}:${Catalog}/ChangeSet/${ResourceId}"},
		{Name: "aws_marketplace_dashboard", Service: "aws-marketplace", Resource: "Dashboard", Template: "arn:${Partition}:aws-marketplace::${Account}:${Catalog}/ReportingData/${FactTable}/Dashboard/${DashboardName}"},
		{Name: "aws_marketplace_deployment_parameter", Service: "aws-marketplace", Resource: "DeploymentParameter", Template: "arn:${Partition}:aws-marketplace:${Region}:${Account}:DeploymentParameter:catalogs/${CatalogName}/products/${ProductId}/${ResourceId}"},
		{Name: "aws_marketplace_entity", Service: "aws-marketplace", Resource: "Entity", Template: "arn:${Partition}:aws-marketplace:${Region}:${Account}:${Catalog}/${EntityType}/${ResourceId}"},
		{Name: "aws_marketplace_invoice_submission_task", Service: "aws-marketplace", Resource: "InvoiceSubmissionTask", Template: "arn:${Partition}:aws-marketplace:${Region}:${Account}:catalog/${Catalog}/invoice-submission-task/${ResourceId}"},
		{Name: "aws_marketplace_issued_tax_invoice", Service: "aws-marketplace", Resource: "IssuedTaxInvoice", Template: "arn:${Partition}:aws-marketplace:${Region}:${Account}:catalog/${Catalog}/issued-tax-invoice/${ResourceId}"},
		{Name: "aws_marketplace_listing", Service: "aws-marketplace", Resource: "Listing", Template: "arn:${Partition}:aws-marketplace:::catalog/${CatalogName}/listing/${ListingId}"},
		{Name: "aws_marketplace_offer", Service: "aws-marketplace", Resource: "Offer", Template: "arn:${Partition}:aws-marketplace:::catalog/${CatalogName}/offer/${OfferId}"},
		{Name: "aws_marketplace_offer_set", Service: "aws-marketplace", Resource: "OfferSet", Template: "arn:${Partition}:aws-marketplace:::catalog/${CatalogName}/offerSet/${OfferSetId}"},
		{Name: "aws_marketplace_product", Service: "aws-marketplace", Resource: "Product", Template: "arn:${Partition}:aws-marketplace:::catalog/${CatalogName}/product/${ProductId}"},
		{Name: "aws_marketplace_purchase_option", Service: "aws-marketplace", Resource: "PurchaseOption", Template: "arn:${Partition}:aws-marketplace:::catalog/${CatalogName}/purchaseOption/${PurchaseOptionId}"},
		{Name: "aws_marketplace_seller_dashboard", Service: "aws-marketplace", Resource: "SellerDashboard", Template: "arn:${Partition}:aws-marketplace::${Account}:${Catalog}/ReportingData/${FactTable}/Dashboard/${DashboardName}"},
		{Name: "aws_marketplace_tax_compliance_profile", Service: "aws-marketplace", Resource: "TaxComplianceProfile", Template: "arn:${Partition}:aws-marketplace:${Region}:${Account}:tax-compliance-profile/${ResourceId}"},
		{Name: "aws_marketplace_tax_compliance_profile_change_task", Service: "aws-marketplace", Resource: "TaxComplianceProfileChangeTask", Template: "arn:${Partition}:aws-marketplace:${Region}:${Account}:tax-compliance-profile-change-task/${ResourceId}"},
		{Name: "aws_marketplace_verification_evidence", Service: "aws-marketplace", Resource: "VerificationEvidence", Template: "arn:${Partition}:aws-marketplace:${Region}:${Account}:verification-type/${VerificationType}/verification-evidence/${ResourceId}"},
	})
}
