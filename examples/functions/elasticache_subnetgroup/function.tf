# arn:aws:elasticache:ap-northeast-1:111111111111:subnetgroup:cache-subnet-group-name
output "elasticache_subnetgroup" {
  value = provider::arn::elasticache_subnetgroup("cache-subnet-group-name")
}
