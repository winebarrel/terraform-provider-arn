# arn:aws:mediastore:ap-northeast-1:111111111111:container/container-name/object-path
output "mediastore_object" {
  value = provider::arn::mediastore_object("container-name", "object-path")
}
