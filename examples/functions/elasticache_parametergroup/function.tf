# arn:aws:elasticache:ap-northeast-1:111111111111:parametergroup:cache-parameter-group-name
output "elasticache_parametergroup" {
  value = provider::arn::elasticache_parametergroup("cache-parameter-group-name")
}
