# arn:aws:cloudfront::111111111111:vpcorigin/id
output "cloudfront_vpcorigin" {
  value = provider::arn::cloudfront_vpcorigin("id")
}
