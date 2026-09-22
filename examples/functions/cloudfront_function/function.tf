# arn:aws:cloudfront::111111111111:function/name
output "cloudfront_function" {
  value = provider::arn::cloudfront_function("name")
}
