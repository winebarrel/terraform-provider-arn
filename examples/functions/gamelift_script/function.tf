# arn:aws:gamelift:ap-northeast-1:111111111111:script/script-id
output "gamelift_script" {
  value = provider::arn::gamelift_script("script-id")
}
