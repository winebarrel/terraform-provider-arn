# arn:aws:iotfleetwise:ap-northeast-1:111111111111:decoder-manifest/name
output "iotfleetwise_decodermanifest" {
  value = provider::arn::iotfleetwise_decodermanifest("name")
}
