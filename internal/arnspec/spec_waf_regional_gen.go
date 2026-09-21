// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: waf-regional
// Source: https://servicereference.us-east-1.amazonaws.com/v1/waf-regional/waf-regional.json
// Functions: 13
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "waf_regional_bytematchset", Service: "waf-regional", Resource: "bytematchset", Template: "arn:${Partition}:waf-regional:${Region}:${Account}:bytematchset/${Id}"},
		{Name: "waf_regional_geomatchset", Service: "waf-regional", Resource: "geomatchset", Template: "arn:${Partition}:waf-regional:${Region}:${Account}:geomatchset/${Id}"},
		{Name: "waf_regional_ipset", Service: "waf-regional", Resource: "ipset", Template: "arn:${Partition}:waf-regional:${Region}:${Account}:ipset/${Id}"},
		{Name: "waf_regional_loadbalancer_app", Service: "waf-regional", Resource: "loadbalancer/app/", Template: "arn:${Partition}:elasticloadbalancing:${Region}:${Account}:loadbalancer/app/${LoadBalancerName}/${LoadBalancerId}"},
		{Name: "waf_regional_ratebasedrule", Service: "waf-regional", Resource: "ratebasedrule", Template: "arn:${Partition}:waf-regional:${Region}:${Account}:ratebasedrule/${Id}"},
		{Name: "waf_regional_regexmatchset", Service: "waf-regional", Resource: "regexmatchset", Template: "arn:${Partition}:waf-regional:${Region}:${Account}:regexmatch/${Id}"},
		{Name: "waf_regional_regexpatternset", Service: "waf-regional", Resource: "regexpatternset", Template: "arn:${Partition}:waf-regional:${Region}:${Account}:regexpatternset/${Id}"},
		{Name: "waf_regional_rule", Service: "waf-regional", Resource: "rule", Template: "arn:${Partition}:waf-regional:${Region}:${Account}:rule/${Id}"},
		{Name: "waf_regional_rulegroup", Service: "waf-regional", Resource: "rulegroup", Template: "arn:${Partition}:waf-regional:${Region}:${Account}:rulegroup/${Id}"},
		{Name: "waf_regional_sizeconstraintset", Service: "waf-regional", Resource: "sizeconstraintset", Template: "arn:${Partition}:waf-regional:${Region}:${Account}:sizeconstraintset/${Id}"},
		{Name: "waf_regional_sqlinjectionmatchset", Service: "waf-regional", Resource: "sqlinjectionmatchset", Template: "arn:${Partition}:waf-regional:${Region}:${Account}:sqlinjectionset/${Id}"},
		{Name: "waf_regional_webacl", Service: "waf-regional", Resource: "webacl", Template: "arn:${Partition}:waf-regional:${Region}:${Account}:webacl/${Id}"},
		{Name: "waf_regional_xssmatchset", Service: "waf-regional", Resource: "xssmatchset", Template: "arn:${Partition}:waf-regional:${Region}:${Account}:xssmatchset/${Id}"},
	})
}
