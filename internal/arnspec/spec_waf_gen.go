// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: waf
// Source: https://servicereference.us-east-1.amazonaws.com/v1/waf/waf.json
// Functions: 12
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "waf_bytematchset", Service: "waf", Resource: "bytematchset", Template: "arn:${Partition}:waf::${Account}:bytematchset/${Id}"},
		{Name: "waf_geomatchset", Service: "waf", Resource: "geomatchset", Template: "arn:${Partition}:waf::${Account}:geomatchset/${Id}"},
		{Name: "waf_ipset", Service: "waf", Resource: "ipset", Template: "arn:${Partition}:waf::${Account}:ipset/${Id}"},
		{Name: "waf_ratebasedrule", Service: "waf", Resource: "ratebasedrule", Template: "arn:${Partition}:waf::${Account}:ratebasedrule/${Id}"},
		{Name: "waf_regexmatchset", Service: "waf", Resource: "regexmatchset", Template: "arn:${Partition}:waf::${Account}:regexmatch/${Id}"},
		{Name: "waf_regexpatternset", Service: "waf", Resource: "regexpatternset", Template: "arn:${Partition}:waf::${Account}:regexpatternset/${Id}"},
		{Name: "waf_rule", Service: "waf", Resource: "rule", Template: "arn:${Partition}:waf::${Account}:rule/${Id}"},
		{Name: "waf_rulegroup", Service: "waf", Resource: "rulegroup", Template: "arn:${Partition}:waf::${Account}:rulegroup/${Id}"},
		{Name: "waf_sizeconstraintset", Service: "waf", Resource: "sizeconstraintset", Template: "arn:${Partition}:waf::${Account}:sizeconstraintset/${Id}"},
		{Name: "waf_sqlinjectionmatchset", Service: "waf", Resource: "sqlinjectionmatchset", Template: "arn:${Partition}:waf::${Account}:sqlinjectionset/${Id}"},
		{Name: "waf_webacl", Service: "waf", Resource: "webacl", Template: "arn:${Partition}:waf::${Account}:webacl/${Id}"},
		{Name: "waf_xssmatchset", Service: "waf", Resource: "xssmatchset", Template: "arn:${Partition}:waf::${Account}:xssmatchset/${Id}"},
	})
}
