# arn:aws:cloudfront::111111111111:key-value-store/resource-id
output "cloudfront_keyvaluestore_key_value_store" {
  value = provider::arn::cloudfront_keyvaluestore_key_value_store("resource-id")
}
