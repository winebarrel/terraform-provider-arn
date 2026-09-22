# arn:aws:mediastore:ap-northeast-1:111111111111:container/container-name/folder-path
output "mediastore_folder" {
  value = provider::arn::mediastore_folder("container-name", "folder-path")
}
