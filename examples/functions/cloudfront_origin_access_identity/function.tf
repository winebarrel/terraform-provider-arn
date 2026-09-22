# arn:aws:cloudfront::111111111111:origin-access-identity/id
output "cloudfront_origin_access_identity" {
  value = provider::arn::cloudfront_origin_access_identity("id")
}
