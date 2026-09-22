# arn:aws:cloudfront::111111111111:anycast-ip-list/id
output "cloudfront_anycast_ip_list" {
  value = provider::arn::cloudfront_anycast_ip_list("id")
}
