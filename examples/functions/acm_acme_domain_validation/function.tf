# arn:aws:acm:ap-northeast-1:111111111111:acme-endpoint/acme-endpoint-id/acme-domain-validation/acme-domain-validation-id
output "acm_acme_domain_validation" {
  value = provider::arn::acm_acme_domain_validation("acme-endpoint-id", "acme-domain-validation-id")
}
