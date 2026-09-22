# arn:aws:mediaconnect:ap-northeast-1:offering:offering-id
output "mediaconnect_offering" {
  value = provider::arn::mediaconnect_offering("offering-id")
}
