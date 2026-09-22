# arn:aws:mediaconnect:ap-northeast-1:111111111111:source:source-id:source-name
output "mediaconnect_source" {
  value = provider::arn::mediaconnect_source("source-id", "source-name")
}
