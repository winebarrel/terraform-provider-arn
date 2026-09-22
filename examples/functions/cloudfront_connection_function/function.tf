# arn:aws:cloudfront::111111111111:connection-function/id
output "cloudfront_connection_function" {
  value = provider::arn::cloudfront_connection_function("id")
}
