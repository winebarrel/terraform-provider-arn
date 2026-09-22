# arn:aws:cloudfront::111111111111:field-level-encryption-config/id
output "cloudfront_field_level_encryption_config" {
  value = provider::arn::cloudfront_field_level_encryption_config("id")
}
