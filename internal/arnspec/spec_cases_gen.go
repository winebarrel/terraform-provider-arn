// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: cases
// Source: https://servicereference.us-east-1.amazonaws.com/v1/cases/cases.json
// Functions: 7
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "cases_case", Service: "cases", Resource: "Case", Template: "arn:${Partition}:cases:${Region}:${Account}:domain/${DomainId}/case/${CaseId}"},
		{Name: "cases_case_rule", Service: "cases", Resource: "CaseRule", Template: "arn:${Partition}:cases:${Region}:${Account}:domain/${DomainId}/case-rule/${CaseRuleId}"},
		{Name: "cases_domain", Service: "cases", Resource: "Domain", Template: "arn:${Partition}:cases:${Region}:${Account}:domain/${DomainId}"},
		{Name: "cases_field", Service: "cases", Resource: "Field", Template: "arn:${Partition}:cases:${Region}:${Account}:domain/${DomainId}/field/${FieldId}"},
		{Name: "cases_layout", Service: "cases", Resource: "Layout", Template: "arn:${Partition}:cases:${Region}:${Account}:domain/${DomainId}/layout/${LayoutId}"},
		{Name: "cases_related_item", Service: "cases", Resource: "RelatedItem", Template: "arn:${Partition}:cases:${Region}:${Account}:domain/${DomainId}/case/${CaseId}/related-item/${RelatedItemId}"},
		{Name: "cases_template", Service: "cases", Resource: "Template", Template: "arn:${Partition}:cases:${Region}:${Account}:domain/${DomainId}/template/${TemplateId}"},
	})
}
