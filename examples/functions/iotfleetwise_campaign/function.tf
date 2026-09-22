# arn:aws:iotfleetwise:ap-northeast-1:111111111111:campaign/campaign-name
output "iotfleetwise_campaign" {
  value = provider::arn::iotfleetwise_campaign("campaign-name")
}
