# arn:aws:acm:ap-northeast-1:111111111111:acme-endpoint/acme-endpoint-id
output "acm_acme_endpoint" {
  value = provider::arn::acm_acme_endpoint("acme-endpoint-id")
}
