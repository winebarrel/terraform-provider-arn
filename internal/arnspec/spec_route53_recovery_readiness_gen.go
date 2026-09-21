// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: route53-recovery-readiness
// Source: https://servicereference.us-east-1.amazonaws.com/v1/route53-recovery-readiness/route53-recovery-readiness.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "route53_recovery_readiness_cell", Service: "route53-recovery-readiness", Resource: "cell", Template: "arn:${Partition}:route53-recovery-readiness::${Account}:cell/${ResourceId}"},
		{Name: "route53_recovery_readiness_readinesscheck", Service: "route53-recovery-readiness", Resource: "readinesscheck", Template: "arn:${Partition}:route53-recovery-readiness::${Account}:readiness-check/${ResourceId}"},
		{Name: "route53_recovery_readiness_recoverygroup", Service: "route53-recovery-readiness", Resource: "recoverygroup", Template: "arn:${Partition}:route53-recovery-readiness::${Account}:recovery-group/${ResourceId}"},
		{Name: "route53_recovery_readiness_resourceset", Service: "route53-recovery-readiness", Resource: "resourceset", Template: "arn:${Partition}:route53-recovery-readiness::${Account}:resource-set/${ResourceId}"},
	})
}
