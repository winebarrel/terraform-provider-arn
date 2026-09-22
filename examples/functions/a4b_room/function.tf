# arn:aws:a4b:ap-northeast-1:111111111111:room/resource-id
output "a4b_room" {
  value = provider::arn::a4b_room("resource-id")
}
