# arn:aws:cloudfront::111111111111:streaming-distribution/distribution-id
output "cloudfront_streaming_distribution" {
  value = provider::arn::cloudfront_streaming_distribution("distribution-id")
}
