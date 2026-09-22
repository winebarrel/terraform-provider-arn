# arn:aws:nimble:ap-northeast-1:111111111111:studio/studio-id
output "nimble_studio" {
  value = provider::arn::nimble_studio("studio-id")
}
