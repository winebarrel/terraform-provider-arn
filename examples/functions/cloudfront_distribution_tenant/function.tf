# arn:aws:cloudfront::111111111111:distribution-tenant/id
output "cloudfront_distribution_tenant" {
  value = provider::arn::cloudfront_distribution_tenant("id")
}
