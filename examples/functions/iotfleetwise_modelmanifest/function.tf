# arn:aws:iotfleetwise:ap-northeast-1:111111111111:model-manifest/name
output "iotfleetwise_modelmanifest" {
  value = provider::arn::iotfleetwise_modelmanifest("name")
}
