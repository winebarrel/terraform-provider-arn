# arn:aws:iotfleetwise:ap-northeast-1:111111111111:state-template/state-template-id
output "iotfleetwise_statetemplate" {
  value = provider::arn::iotfleetwise_statetemplate("state-template-id")
}
