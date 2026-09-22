# arn:aws:acm:ap-northeast-1:111111111111:acme-endpoint/acme-endpoint-id/acme-external-account-binding/external-account-binding-id
output "acm_acme_external_account_binding" {
  value = provider::arn::acm_acme_external_account_binding("acme-endpoint-id", "external-account-binding-id")
}
