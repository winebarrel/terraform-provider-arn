// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: route53-recovery-control-config
// Source: https://servicereference.us-east-1.amazonaws.com/v1/route53-recovery-control-config/route53-recovery-control-config.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "route53_recovery_control_config_cluster", Service: "route53-recovery-control-config", Resource: "cluster", Template: "arn:${Partition}:route53-recovery-control::${Account}:cluster/${ResourceId}"},
		{Name: "route53_recovery_control_config_controlpanel", Service: "route53-recovery-control-config", Resource: "controlpanel", Template: "arn:${Partition}:route53-recovery-control::${Account}:controlpanel/${ControlPanelId}"},
		{Name: "route53_recovery_control_config_routingcontrol", Service: "route53-recovery-control-config", Resource: "routingcontrol", Template: "arn:${Partition}:route53-recovery-control::${Account}:controlpanel/${ControlPanelId}/routingcontrol/${RoutingControlId}"},
		{Name: "route53_recovery_control_config_safetyrule", Service: "route53-recovery-control-config", Resource: "safetyrule", Template: "arn:${Partition}:route53-recovery-control::${Account}:controlpanel/${ControlPanelId}/safetyrule/${SafetyRuleId}"},
	})
}
