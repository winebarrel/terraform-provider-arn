// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: acm
// Source: https://servicereference.us-east-1.amazonaws.com/v1/acm/acm.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "acm_acme_domain_validation", Service: "acm", Resource: "acme-domain-validation", Template: "arn:${Partition}:acm:${Region}:${Account}:acme-endpoint/${AcmeEndpointId}/acme-domain-validation/${AcmeDomainValidationId}"},
		{Name: "acm_acme_endpoint", Service: "acm", Resource: "acme-endpoint", Template: "arn:${Partition}:acm:${Region}:${Account}:acme-endpoint/${AcmeEndpointId}"},
		{Name: "acm_acme_external_account_binding", Service: "acm", Resource: "acme-external-account-binding", Template: "arn:${Partition}:acm:${Region}:${Account}:acme-endpoint/${AcmeEndpointId}/acme-external-account-binding/${ExternalAccountBindingId}"},
		{Name: "acm_certificate", Service: "acm", Resource: "certificate", Template: "arn:${Partition}:acm:${Region}:${Account}:certificate/${CertificateId}"},
	})
}
