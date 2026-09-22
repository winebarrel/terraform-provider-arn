# arn:aws:cloudfront::111111111111:trust-store/id
output "cloudfront_trust_store" {
  value = provider::arn::cloudfront_trust_store("id")
}
