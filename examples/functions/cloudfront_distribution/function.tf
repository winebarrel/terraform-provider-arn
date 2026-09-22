# arn:aws:cloudfront::111111111111:distribution/distribution-id
output "cloudfront_distribution" {
  value = provider::arn::cloudfront_distribution("distribution-id")
}
