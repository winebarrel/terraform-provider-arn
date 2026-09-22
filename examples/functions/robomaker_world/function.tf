# arn:aws:robomaker:ap-northeast-1:111111111111:world/world-id
output "robomaker_world" {
  value = provider::arn::robomaker_world("world-id")
}
