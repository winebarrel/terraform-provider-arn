# arn:aws:mediastore:ap-northeast-1:111111111111:container/container-name
output "mediastore_container" {
  value = provider::arn::mediastore_container("container-name")
}
