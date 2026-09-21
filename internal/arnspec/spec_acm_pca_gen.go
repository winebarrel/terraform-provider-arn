// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: acm-pca
// Source: https://servicereference.us-east-1.amazonaws.com/v1/acm-pca/acm-pca.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "acm_pca_certificate_authority", Service: "acm-pca", Resource: "certificate-authority", Template: "arn:${Partition}:acm-pca:${Region}:${Account}:certificate-authority/${CertificateAuthorityId}"},
	})
}
