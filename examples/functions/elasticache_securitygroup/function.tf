# arn:aws:elasticache:ap-northeast-1:111111111111:securitygroup:cache-security-group-name
output "elasticache_securitygroup" {
  value = provider::arn::elasticache_securitygroup("cache-security-group-name")
}
