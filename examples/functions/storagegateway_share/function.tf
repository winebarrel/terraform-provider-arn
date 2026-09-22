# arn:aws:storagegateway:ap-northeast-1:111111111111:share/share-id
output "storagegateway_share" {
  value = provider::arn::storagegateway_share("share-id")
}
