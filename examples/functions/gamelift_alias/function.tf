# arn:aws:gamelift:ap-northeast-1::alias/alias-id
output "gamelift_alias" {
  value = provider::arn::gamelift_alias("alias-id")
}
