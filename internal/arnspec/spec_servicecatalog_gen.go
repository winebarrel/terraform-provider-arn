// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: servicecatalog
// Source: https://servicereference.us-east-1.amazonaws.com/v1/servicecatalog/servicecatalog.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "servicecatalog_application", Service: "servicecatalog", Resource: "Application", Template: "arn:${Partition}:servicecatalog:${Region}:${Account}:/applications/${ApplicationId}"},
		{Name: "servicecatalog_attribute_group", Service: "servicecatalog", Resource: "AttributeGroup", Template: "arn:${Partition}:servicecatalog:${Region}:${Account}:/attribute-groups/${AttributeGroupId}"},
		{Name: "servicecatalog_portfolio", Service: "servicecatalog", Resource: "Portfolio", Template: "arn:${Partition}:catalog:${Region}:${Account}:portfolio/${PortfolioId}"},
		{Name: "servicecatalog_product", Service: "servicecatalog", Resource: "Product", Template: "arn:${Partition}:catalog:${Region}:${Account}:product/${ProductId}"},
	})
}
