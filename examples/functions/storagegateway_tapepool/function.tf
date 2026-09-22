# arn:aws:storagegateway:ap-northeast-1:111111111111:tapepool/pool-id
output "storagegateway_tapepool" {
  value = provider::arn::storagegateway_tapepool("pool-id")
}
