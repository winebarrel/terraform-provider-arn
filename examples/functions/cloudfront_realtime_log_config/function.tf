# arn:aws:cloudfront::111111111111:realtime-log-config/name
output "cloudfront_realtime_log_config" {
  value = provider::arn::cloudfront_realtime_log_config("name")
}
