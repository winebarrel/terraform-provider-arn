# arn:aws:iotfleetwise:ap-northeast-1:111111111111:signal-catalog/name
output "iotfleetwise_signalcatalog" {
  value = provider::arn::iotfleetwise_signalcatalog("name")
}
