# arn:aws:elasticache:ap-northeast-1:111111111111:usergroup:user-group-id
output "elasticache_usergroup" {
  value = provider::arn::elasticache_usergroup("user-group-id")
}
