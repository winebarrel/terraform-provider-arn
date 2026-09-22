# arn:aws:iot:ap-northeast-1:111111111111:mitigationaction/mitigation-action-name
output "iot_mitigationaction" {
  value = provider::arn::iot_mitigationaction("mitigation-action-name")
}
