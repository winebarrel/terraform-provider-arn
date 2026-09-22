# arn:aws:mediaconnect:ap-northeast-1:111111111111:bridge:bridge-id:bridge-name
output "mediaconnect_bridge" {
  value = provider::arn::mediaconnect_bridge("bridge-id", "bridge-name")
}
