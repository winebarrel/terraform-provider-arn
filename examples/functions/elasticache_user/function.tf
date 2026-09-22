# arn:aws:elasticache:ap-northeast-1:111111111111:user:user-id
output "elasticache_user" {
  value = provider::arn::elasticache_user("user-id")
}
