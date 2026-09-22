# arn:aws:gamelift:ap-northeast-1:111111111111:build/build-id
output "gamelift_build" {
  value = provider::arn::gamelift_build("build-id")
}
