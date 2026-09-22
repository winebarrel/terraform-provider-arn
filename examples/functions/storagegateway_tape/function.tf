# arn:aws:storagegateway:ap-northeast-1:111111111111:tape/tape-barcode
output "storagegateway_tape" {
  value = provider::arn::storagegateway_tape("tape-barcode")
}
