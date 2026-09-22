# arn:aws:cloudfront::111111111111:cache-policy/id
output "cloudfront_cache_policy" {
  value = provider::arn::cloudfront_cache_policy("id")
}
