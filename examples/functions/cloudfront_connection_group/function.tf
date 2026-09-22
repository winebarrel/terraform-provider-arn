# arn:aws:cloudfront::111111111111:connection-group/id
output "cloudfront_connection_group" {
  value = provider::arn::cloudfront_connection_group("id")
}
