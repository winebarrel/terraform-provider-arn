# arn:aws:mediaconnect:ap-northeast-1:111111111111:flow:flow-id:flow-name
output "mediaconnect_flow" {
  value = provider::arn::mediaconnect_flow("flow-id", "flow-name")
}
