# arn:aws:cloudfront::111111111111:origin-request-policy/id
output "cloudfront_origin_request_policy" {
  value = provider::arn::cloudfront_origin_request_policy("id")
}
