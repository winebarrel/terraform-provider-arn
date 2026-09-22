# arn:aws:storagegateway:ap-northeast-1:111111111111:fs-association/fsa-id
output "storagegateway_fs_association" {
  value = provider::arn::storagegateway_fs_association("fsa-id")
}
