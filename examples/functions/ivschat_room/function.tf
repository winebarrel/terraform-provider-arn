# arn:aws:ivschat:ap-northeast-1:111111111111:room/resource-id
output "ivschat_room" {
  value = provider::arn::ivschat_room("resource-id")
}
