# arn:aws:cloudfront::111111111111:response-headers-policy/id
output "cloudfront_response_headers_policy" {
  value = provider::arn::cloudfront_response_headers_policy("id")
}
