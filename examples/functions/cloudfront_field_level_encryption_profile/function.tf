# arn:aws:cloudfront::111111111111:field-level-encryption-profile/id
output "cloudfront_field_level_encryption_profile" {
  value = provider::arn::cloudfront_field_level_encryption_profile("id")
}
