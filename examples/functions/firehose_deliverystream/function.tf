# arn:aws:firehose:ap-northeast-1:111111111111:deliverystream/delivery-stream-name
output "firehose_deliverystream" {
  value = provider::arn::firehose_deliverystream("delivery-stream-name")
}
