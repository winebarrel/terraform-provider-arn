# arn:aws:iot:ap-northeast-1:111111111111:otaupdate/ota-update-id
output "iot_otaupdate" {
  value = provider::arn::iot_otaupdate("ota-update-id")
}
