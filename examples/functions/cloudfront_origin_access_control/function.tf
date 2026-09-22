# arn:aws:cloudfront::111111111111:origin-access-control/id
output "cloudfront_origin_access_control" {
  value = provider::arn::cloudfront_origin_access_control("id")
}
