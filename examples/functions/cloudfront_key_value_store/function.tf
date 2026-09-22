# arn:aws:cloudfront::111111111111:key-value-store/name
output "cloudfront_key_value_store" {
  value = provider::arn::cloudfront_key_value_store("name")
}
